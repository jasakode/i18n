package i18n_test

import (
	"fmt"
	"testing"

	"github.com/jasakode/i18n"
)

func TestXxx(t *testing.T) {
	var locale i18n.Locale
	fmt.Println(locale.Parse("en;id"))

	dict := i18n.New(map[i18n.Locale]string{}, i18n.Options{
		Setup: func(app *i18n.App) i18n.Configuration {
			return i18n.Configuration{
				OnTranslate: func(text string, from, to i18n.Locale) (result string, err error) {
					fmt.Println("Hallo", app)
					return
				},
			}
		},
		// Setup: func(config i18n.Configuration) {
		// 	config.OnTranslate = func(text string, from, to i18n.Locale) (result string, err error) {
		// 		return
		// 	}
		// 	config.OnErrorFallback = func(text string, locale i18n.Locale) string {
		// 		return ""
		// 	}
		// },
	})

	fmt.Println(dict.T("", i18n.LocaleAR))
}
