// Package i18n provides a lightweight, high-performance internationalization (i18n)
// solution for Go applications, inspired by the elegant developer experience of next-intl.
// It leverages Go's native embed.FS to bundle translation files into a single binary.
package i18n

import (
	"fmt"
	"sync"
)

type App struct {
	mu         sync.RWMutex
	dictionary map[Locale]map[string]any

	// this is default locale
	LOCALE Locale
	// default timezone
	TIMEZONE TimeZone
}

type Configuration struct {
	//
	OnTranslate func(text string, from Locale, to Locale) (result string, err error)
	// for T(text string, locale Locale) method
	OnErrorFallback func(text string, locale Locale) string
}

type Options struct {
	DefaultLocale   Locale
	DefaultTimezone TimeZone
	Setup           func(app *App) Configuration
}

func New(dictionary map[Locale]string, options ...Options) *App {
	app := App{}
	if len(options) > 0 && options[0].Setup != nil {
		options[0].Setup(&app)
	}
	return &App{}
}

func (app *App) T(text string, locale Locale) string {
	app.mu.Lock()
	defer app.mu.Unlock()

	fmt.Println(app.dictionary)
	return text
}
