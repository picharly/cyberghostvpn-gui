package locales

import (
	"cyberghostvpn-gui/logger"
	"embed"
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/tailscale/hujson"
	"golang.org/x/text/language"

	_ "embed"
)

// Language files

//go:embed *.jsonc
var langFS embed.FS

var availableLanguages = []lang{}
var bundle *i18n.Bundle
var currentLanguage lang
var defaultLanguage = language.English
var loc *i18n.Localizer

type lang struct {
	Name       string
	Code       string
	Dictionary []byte
	Tag        language.Tag
}

type Message struct {
	ID    string `json:"id"`
	Other string `json:"other"`
}

type Data struct {
	Messages []Message `json:"messages"`
}

type Variable struct {
	Name  string
	Value interface{}
}

// GetCurrentLocale returns the current language in use.
func GetCurrentLocale() lang {
	return currentLanguage
}

// GetLocales returns a list of all available languages that can be used with the locales.Init() function.
func GetLocales() []lang {
	return availableLanguages
}

// GetSystemLocale returns the current system locale, first by reading the LANG,
// then LC_ALL and finally LC_CTYPE environment variables. If none of those
// variables are set, the OS name is returned. If the returned locale ends with
// a dot, it is removed.
func GetSystemLocale() string {
	locale := os.Getenv("LANG")
	if locale == "" {
		locale = os.Getenv("LC_ALL")
	}
	if locale == "" {
		locale = os.Getenv("LC_CTYPE")
	}
	if locale == "" {
		locale = runtime.GOOS
	}

	idx := strings.IndexFunc(locale, func(r rune) bool {
		return !strings.ContainsRune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ", r)
	})
	if idx != -1 {
		locale = locale[:idx]
	}

	// if strings.Contains(locale, ".") {
	// 	locale = strings.Split(locale, ".")[0]
	// }

	return locale
}

// Init loads the default/configured language and starts the translation trigger. If the locale string is empty, it will
// use the system locale. If the locale string is not empty, it will load the language with the given locale. If the
// locale string is not found, it will load the default language.
func Init(locale string) {
	if len(availableLanguages) == 0 {
		// Load all available languages
		loadLocales()
	}

	// Load the default/configured language
	if len(locale) > 0 {
		load(getLanguageTag(locale))
		setCurrentLocale(locale)
	} else {
		load(getLanguageTag(GetSystemLocale()))
		setCurrentLocale(GetSystemLocale())
	}

	// Start the translation trigger
	GetTrigger().Start()
}

// Get language tag from string name
func getLanguageTag(name string) language.Tag {
	setCurrentLocale(name)
	tag, err := language.Parse(name)
	if err != nil {
		logger.Warnf("cannot load locale '%s': %v", name, err)
		setCurrentLocale("en")
		return language.English
	}
	return tag
}

// load a language in memory
func load(language language.Tag) {
	bundle = i18n.NewBundle(language)
	if err := loadLocale(language); err != nil {
		logger.Errorf("cannot load locale %v: %v", language, err)
		return
	}
	loc = i18n.NewLocalizer(bundle, language.String())
}

// loadLocales loads all available languages
//
// It iterates over the language files in the embedded FS and loads their
// content into memory. It then parses the JSON content of each file and
// extracts the language name and code. If both fields are present, the
// language is added to the list of available languages.
//
// It is called by init() and the list of available languages is used by
// the load() function to determine which language to load.
func loadLocales() {
	files, err := langFS.ReadDir(".")
	if err != nil {
		return
	}

	// Iterate over the files and print their names
	availableLanguages = make([]lang, 0)
	for _, file := range files {
		if !file.IsDir() {
			jsonData, err := fs.ReadFile(langFS, filepath.Join(".", file.Name()))
			if err != nil {
				continue
			}
			newLang := lang{}
			stdJson, err := standardizeJSON(jsonData)
			if err != nil {
				continue
			}

			var data Data
			err = json.Unmarshal(stdJson, &data)
			if err != nil {
				continue
			}

			for _, message := range data.Messages {
				if message.ID == "name" {
					newLang.Name = message.Other
				} else if message.ID == "code" {
					newLang.Code = message.Other
				}
			}

			if len(newLang.Code) > 0 && len(newLang.Name) > 0 {
				newLang.Dictionary = stdJson
			}

			if len(newLang.Name) > 0 {
				availableLanguages = append(availableLanguages, newLang)
			}
		}
	}
}

// newVariable creates a new Variable from a name and a value.
// The variable will be used to replace placeholders in localized strings.
func newVariable(name string, value interface{}) Variable {
	return Variable{Name: name, Value: value}
}

// setCurrentLocale sets the current language from the given locale string. If the
// locale matches the code of a loaded language, it is set as the current language.
// If the locale does not match the code of a loaded language, it is assumed to be
// the name of a language and the locale is set based on the first matching language.
// If no matching language is found, the current language is not changed.
func setCurrentLocale(locale string) {
	for _, l := range availableLanguages {
		if strings.EqualFold(l.Code, locale) || strings.EqualFold(l.Name+"_", locale) {
			currentLanguage = l
			break
		}
	}
}

// loadLocale loads a language in memory
// If the language is not loaded, returns an error
// The language is identified by its locale string
// If the language is loaded, the current language is set to the loaded language
func loadLocale(lang language.Tag) error {
	if bundle != nil {
		var data []byte
		for _, l := range availableLanguages {
			if l.Code == lang.String() {
				data = l.Dictionary
				break
			}
		}

		if _, err := bundle.ParseMessageFileBytes(data, fmt.Sprintf("msg.%s.json", lang.String())); err != nil {
			return fmt.Errorf("cannot load locale %s: %v", lang.String(), err)
		}
	}
	return nil
}

// Standardise JSON to remove comments
func standardizeJSON(b []byte) ([]byte, error) {
	ast, err := hujson.Parse(b)
	if err != nil {
		return b, err
	}
	ast.Standardize()
	return ast.Pack(), nil
}
