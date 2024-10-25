package locales

import (
	"fmt"

	"github.com/nicksnyder/go-i18n/v2/i18n"
)

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
