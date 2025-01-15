#!/bin/bash

# Installation script for Cyberghost VPN - Works on OpenSUSE Tumbleweed
# For now only openvpn works with this client.
#
# Extremely inspired by https://aur.archlinux.org/packages/cyberghostvpn
#
# Needs to be cleaned up and refactored

pkgname=cyberghostvpn
pkgver=1.4.1
pkgrel=12
pkgdesc="CyberGhost VPN"
url="https://www.cyberghostvpn.com"
_variant=fedora-32 #ubuntu-20.04 #
source=(	"https://download.cyberghostvpn.com/linux/cyberghostvpn-${_variant}-${pkgver}.zip"
		      "http://crt.sectigo.com/SectigoRSAOrganizationValidationSecureServerCA.crt"
)

# List of dependencies
depends=(
  bash
  curl
  openvpn
  wireguard-tools
  openresolv
  ca-certificates
  openssl
  zip
)

# Missing dependencies - must be empty
missing=()

# Temporary folder
tmpFolder="/tmp/aur-cyberghost"

# Colors used for terminal ouput
red='\033[0;31m'
green='\033[0;32m'
normal='\033[0m'

# Check if a package has already been installed or not
check_package() {
    if ! [ "$1" == "" ]; then
        result=$(zypper se -xi $1 > /dev/null && echo $?)
        if ! [ "$result" == "0" ]; then
            missing+=( $1 )
            return 1
        fi
    fi
    return 0
}

download_file() {
  if ! [ "$1" == "" ]; then
    echo "Downloading '$1'..."
    curl -O $1
    if [ $? -ne 0 ]; then
      echo "Failed to download $1"
      exit 2
    fi
  fi
}

# Install packages marked as 'missing'
install_missing_packages() {
    if [ ${#missing[@]} -gt 0 ]; then
        echo "Need to install ${#missing[@]} missing package(s):"
        for package in ${missing[@]}
        do
            echo "    - $package"
        done
        echo ""
        reply=y
        read -p "Do you want to continue? [y/n]: " reply
        case $reply in
            [Yy])
                echo -e "Installing missing dependencies...\n"
                sudo zypper in "${missing[@]}"
                ;;
            [Nn])
                echo "Cannot continue without these dependencies."
                exit 3
                ;;
            *)
                install_missing_packages
                ;;
        esac
    fi
}

_archive="${pkgname}-${_variant}-${pkgver}"

prepare() {
	# workaround: build certificate to connect to wireguard servers
	# remove as soon as the certificate is provied by the package
	_wireguard_certificate_servername="washington-s403-i01.cg-dialup.net"
	_wireguard_certificate_server="102.165.48.72:1337"
	true | openssl s_client -verify 5 -connect ${_wireguard_certificate_server} -servername ${_wireguard_certificate_servername}| openssl x509 > "${tmpFolder}/cg-dialup-net.pem"

	sha256sum "${tmpFolder}/cg-dialup-net.pem"
	sha256sum --check <( echo "e878bc02a7fff67ec7dd39242118ea72e5eb6db6c049b4eaae49a3095a054233	${tmpFolder}/cg-dialup-net.pem"  )

	openssl x509 -in "${tmpFolder}/cg-dialup-net.pem" > "${tmpFolder}/wireguard_ca.crt"

  # Extra step for OpenSuSE
  ln -sf /var/lib/ca-certificates/ca-bundle.pem /etc/ssl/certs/ca-certificates.crt
  if ! [ -f /etc/pki/tls/certs/ ]; then
    mkdir -p /etc/pki/tls/certs
  fi
  ln -sf /var/lib/ca-certificates/ca-bundle.pem /etc/pki/tls/certs/ca-bundle.crt
}

package() {
  _installdir=/usr/local/cyberghost
  install -Dm 755 openvpn_wrapper "${_installdir}/wrapper/openvpn_wrapper"
  install -Dm 755 cyberghostvpn_wrapper "${_installdir}/wrapper/cyberghostvpn_wrapper"
  ln -sf "${_installdir}/wrapper/openvpn_wrapper" "${_installdir}/openvpn"

  install -Dm 644 "${tmpFolder}/wireguard_ca.crt" "${_installdir}/certs/wireguard/ca.crt"

  cd "$_archive"

  install -Dm 755 cyberghost/cyberghostvpn "${_installdir}/cyberghostvpn"
  install -Dm 755 cyberghost/update-systemd-resolved "${_installdir}/update-systemd-resolved"

  install -Dm 644 cyberghost/certs/openvpn/ca.crt "${_installdir}/certs/openvpn/ca.crt"
  install -Dm 644 cyberghost/certs/openvpn/client.crt "${_installdir}/certs/openvpn/client.crt"
  install -Dm 644 cyberghost/certs/openvpn/client.key "${_installdir}/certs/openvpn/client.key"

  install -dm 755 usr/bin
  ln -sf ${_installdir}/wrapper/cyberghostvpn_wrapper /usr/bin/cyberghostvpn
  ln -sf ${_installdir}/update-systemd-resolved /usr/bin/update-systemd-resolved
}

if ! [[ $EUID -eq 0 ]]; then
    echo "This script must be run as root."
    exit 1
fi


# Checking is all dependencies hve been installed
echo "Checking dependencies:"
for package in ${depends[@]}
do
    echo -ne "- $package: "
    check_package $package
    if [ "$?" -eq 0 ]; then
        echo -e "${green}installed${normal}"
    else
        echo -e "${red}missing${normal}"
    fi
done

# Install missing packages (if needed)
install_missing_packages

# Check temporary folder
if ! [ -d "$tmpFolder" ]; then
    mkdir -p "$tmpFolder"
fi
cd "$tmpFolder"
rm -fr *


# Downloads files
for file in ${source[@]}
do
  download_file $file
done

# Create wrappers
echo "#!/bin/bash

# put location of openvpn wrapper first in $PATH
# to ensure that cyberghost is calling the openvpn wrapper
export PATH=/usr/local/cyberghost:$PATH

# cyberghostvpn expects wireguard certificate to be located at '../certs/wireguard/ca.crt'
# cd into /usr/local/cyberghostvpn/certs so that pinned certificate will be found
cd /usr/local/cyberghost/certs

/usr/local/cyberghost/cyberghostvpn \"\$@\"
" > $tmpFolder/cyberghostvpn_wrapper

echo "#!/bin/bash

# strip --ncp-disable parameter from call
#  --ncp-disable is removed from openvpn >=2.6
args=(\"\$@\")
for ((i=0;i<\${#args[@]};i++))
do
	if [ \"\${args[\$i]}\" == \"--ncp-disable\" ]
	then
		unset args[\$i]
	fi
done

echo \"\$0: openvpn wrapper for cyberghostvpn utility\"
echo now executing /sbin/openvpn \"\${args[@]}\"
/sbin/openvpn \"\${args[@]}\"
" > $tmpFolder/openvpn_wrapper

# Extracting CyberGhost archive
echo "Extracting archive..."
unzip -o cyberghostvpn-${_variant}-${pkgver}.zip

# Wireguard workaround
prepare

# Install files
echo "Installing files..."
package

# Sudo workaround
echo ""
echo "If Cyberghost connection failed, you need to edit your /etc/sudoers file (or execute 'sudo visudo')"
echo "to change secure_path to include /usr/local/cyberghost:"
echo ""
echo "Ex: Defaults secure_path=\"/usr/local/cyberghost:/usr/sbin:/usr/bin:/sbin:/bin\""
echo ""