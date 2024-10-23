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

//go:embed en.jsonc
var locale_en []byte

// //go:embed fr.jsonc
var locale_fr []byte

var availableLanguages = []lang{}
var currentLanguage lang
var defaultLanguage = language.English
var bundle *i18n.Bundle
var loc *i18n.Localizer

type lang struct {
	Name string
	Code string
	Tag  language.Tag
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

func GetCurrentLocale() lang {
	return currentLanguage
}

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

	if strings.Contains(locale, ".") {
		locale = strings.Split(locale, ".")[0]
	}

	return locale
}

func Init(locale string) {
	if loc == nil {
		// Load all available languages
		loadLocales()

		// Load the default/configured language
		if len(locale) > 0 {
			load(getLanguageTag(locale))
			setCurrentLocale(locale)
		} else {
			load(getLanguageTag(GetSystemLocale()))
			setCurrentLocale(GetSystemLocale())
		}
	}
}

// Get language tag from string name
func getLanguageTag(name string) language.Tag {
	setCurrentLocale(name)
	switch name {
	case "en", "en_GB", "en_US":
		return language.English
	// case "fr", "fr_FR":
	// 	return language.French
	default:
		logger.Warnf("cannot load locale '%s': unavailable language", name)
		setCurrentLocale("en")
		return language.English
	}
}

// load a language in memory
func load(language language.Tag) {
	if bundle == nil {
		bundle = i18n.NewBundle(language)
	}
	if err := loadLocale(language); err != nil {
		logger.Errorf("cannot load locale %v: %v", language, err)
		return
	}
	loc = i18n.NewLocalizer(bundle, language.String())
	// logger.Infof("loaded locale %v", language)
}

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
			f := lang{}
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
					f.Name = message.Other
				} else if message.ID == "code" {
					f.Code = message.Other
				}
			}

			if len(f.Name) > 0 {
				availableLanguages = append(availableLanguages, f)
			}
		}
	}
}

// newVariable creates a new Variable from a name and a value.
// The variable will be used to replace placeholders in localized strings.
func newVariable(name string, value interface{}) Variable {
	return Variable{Name: name, Value: value}
}

func setCurrentLocale(locale string) {
	for _, l := range availableLanguages {
		if strings.EqualFold(l.Code, locale) || strings.EqualFold(l.Name+"_", locale) {
			currentLanguage = l
			break
		}
	}
}

// Text returns the translated string for the given id and optional variables.
// The result is a string in the language previously loaded with the Load function.
// If the language is not loaded, the default language is used.
// If the id is not found, the string "{id}" is returned.
// If a variable is not found, the string "{variableName}" is returned.
func Text(id string, variables ...Variable) string {
	return _text(id, 0, variables...)
}

// TextPlural returns the translated string for the given id and optional variables, using the
// plural form based on the given pluralCount.
// The result is a string in the language previously loaded with the Load function.
// If the language is not loaded, the default language is used.
// If the id is not found, the string "{id}" is returned.
// If a variable is not found, the string "{variableName}" is returned.
func TextPlural(id string, pluralCount int, variables ...Variable) string {
	return _text(id, pluralCount, variables...)
}

// _text returns the translated string for the given id and optional variables, using the
// plural form based on the given pluralCount.
// The result is a string in the language previously loaded with the Load function.
// If the language is not loaded, the default language is used.
// If the id is not found, the string "{id}" is returned.
// If a variable is not found, the string "{variableName}" is returned.
func _text(id string, pluralCount int, variables ...Variable) string {
	if loc == nil {
		load(defaultLanguage)
	}
	if len(id) < 1 {
		return "{EMPTY ID}"
	}
	id = "messages." + id
	var mapVar map[string]interface{}
	if len(variables) > 0 {
		mapVar = make(map[string]interface{})
		for _, variable := range variables {
			if len(variable.Name) > 0 && variable.Value != nil {
				mapVar[variable.Name] = variable.Value
			}
		}
	}
	var text string
	var err error
	if pluralCount > 1 {
		text, err = loc.Localize(&i18n.LocalizeConfig{MessageID: id, TemplateData: mapVar, PluralCount: pluralCount})
	} else {
		text, err = loc.Localize(&i18n.LocalizeConfig{MessageID: id, TemplateData: mapVar})
	}
	if err != nil {
		return fmt.Sprintf("{%s - %v}", id, err)
	} else {
		return text
	}
}

// Load a locale from a JSON constant
func loadLocale(lang language.Tag) error {
	if bundle != nil {
		var data []byte
		var err error
		switch lang {
		case language.AmericanEnglish, language.English, language.BritishEnglish:
			data, err = standardizeJSON(locale_en)
			// data, err = standardizeJSON([]byte(locale_enUS))
		case language.French:
			data, err = standardizeJSON(locale_fr)
		}
		if err == nil {
			if _, err := bundle.ParseMessageFileBytes(data, fmt.Sprintf("msg.%s.json", lang.String())); err != nil {
				logger.Fatalf("cannot load locale %s: %v\n", lang.String(), err)
			}

		} else {
			return fmt.Errorf("cannot read data for language %s: %v", lang.String(), err)
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
