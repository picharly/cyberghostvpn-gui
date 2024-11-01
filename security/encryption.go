package security

/*
	Encryption/Decryption tool
	1. Based on samples found on Internet (see comments below)
	2. Use a main key that should be changed for each project to use a different encryption
*/

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"cyberghostvpn-gui/locales"
	"encoding/hex"
	"fmt"
	"io"
)

// Key used for crypting
var _keyString = "7B23DB33BD04F76FF776DF2326D28F861C731A56E5D3526720955661D24EFC3C"
var encryptionPassword = ""

// Encrypt encrypts a string using AES encryption with the global key set by SetEncryptionPassword (default is empty string)
// and returns the encrypted string as a hexadecimal string. If an error occurs, it returns an empty string and the error.
func Encrypt(stringToEncrypt string) (string, error) {
	return encrypt(stringToEncrypt, _keyString+encryptionPassword)
}

// EncryptPassword encrypts a string using AES encryption with the main key set during the compilation
// (using the -ldflags "-X main._keyString=<new key>" flag) and returns the encrypted string as a hexadecimal string.
// If an error occurs, it returns an empty string and the error.
func EncryptPassword(stringToEncrypt string) (string, error) {
	return encrypt(stringToEncrypt, _keyString)
}

// Decrypt decrypts a string using AES encryption with the global key set by SetEncryptionPassword (default is empty string)
// and returns the decrypted string and an error if any. If an error occurs, it returns an empty string and the error.
func Decrypt(encryptedString string) (string, error) {
	return decrypt(encryptedString, _keyString+encryptionPassword)
}

// DecryptPassword decrypts a string using AES encryption with the main key set during the compilation
// (using the -ldflags "-X main._keyString=<new key>" flag) and returns the decrypted string and an error if any.
// If an error occurs, it returns an empty string and the error.
func DecryptPassword(encryptedString string) (string, error) {
	return decrypt(encryptedString, _keyString)
}

// SetEncryptionPassword sets the password to be used for encryption and decryption.
// The password is used to generate a key that is used for AES encryption and decryption.
// If the password is not set, the functions will use an empty string as the key.
func SetEncryptionPassword(password string) {
	encryptionPassword = password
}

// Encrypt : encrypt text to AES
func encrypt(stringToEncrypt string, keyString string) (string, error) {

	//Since the key is in string, we need to convert decode it to bytes
	key, _ := hex.DecodeString(keyString)
	plaintext := []byte(stringToEncrypt)

	//Create a new Cipher Block from the key
	block, err := aes.NewCipher(key)
	if err != nil {
		return locales.Text("err.sec0"), err
	}

	//Create a new GCM - https://en.wikipedia.org/wiki/Galois/Counter_Mode
	//https://golang.org/pkg/crypto/cipher/#NewGCM
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return locales.Text("err.sec1"), err
	}

	//Create a nonce. Nonce should be from GCM
	nonce := make([]byte, aesGCM.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return locales.Text("err.sec2"), err
	}

	//Encrypt the data using aesGCM.Seal
	//Since we don't want to save the nonce somewhere else in this case, we add it as a prefix to the encrypted data. The first nonce argument in Seal is the prefix.
	ciphertext := aesGCM.Seal(nonce, nonce, plaintext, nil)
	return fmt.Sprintf("%x", ciphertext), nil

	// Function from https://www.melvinvivas.com/how-to-encrypt-and-decrypt-data-using-aes
}

// Decrypt : decrypt text to AES
func decrypt(encryptedString string, keyString string) (string, error) {
	key, _ := hex.DecodeString(keyString)
	enc, _ := hex.DecodeString(encryptedString)

	//Create a new Cipher Block from the key
	block, err := aes.NewCipher(key)
	if err != nil {
		return locales.Text("err.sec3"), err
	}

	//Create a new GCM
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return locales.Text("err.sec4"), err
	}

	//Get the nonce size
	nonceSize := aesGCM.NonceSize()

	//Extract the nonce from the encrypted data
	if len(enc) < nonceSize {
		return locales.Text("err.sec5"), fmt.Errorf(locales.Text("err.sec6"))
	}
	nonce, ciphertext := enc[:nonceSize], enc[nonceSize:]

	//Decrypt the data
	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "error", err
	}

	return string(plaintext), nil

	// Function from https://www.melvinvivas.com/how-to-encrypt-and-decrypt-data-using-aes
}

// GenerateNewRandomKey could help in generating a new encryption key
func GenerateNewRandomKey() string {
	bytes := make([]byte, 32) //generate a random 32 byte key for AES-256
	if _, err := rand.Read(bytes); err != nil {
		panic(err.Error())
	}

	key := hex.EncodeToString(bytes)
	return key
}
