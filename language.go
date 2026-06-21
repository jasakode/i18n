// Package i18n provides a lightweight, high-performance internationalization (i18n)
// solution for Go applications, inspired by the elegant developer experience of next-intl.
// It leverages Go's native embed.FS to bundle translation files into a single binary.
package i18n

var LanguageAR Language = Language{
	ID:          LocaleAR,
	Flag:        "🇸🇦",
	Direction:   DirectionLTR,
	EnglishName: "Arabic",
	NativeName:  "العربية",
}

var Languages []Language = []Language{
	LanguageAR,
}
