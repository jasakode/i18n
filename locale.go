// Package i18n provides a lightweight, high-performance internationalization (i18n)
// solution for Go applications, inspired by the elegant developer experience of next-intl.
// It leverages Go's native embed.FS to bundle translation files into a single binary.
package i18n

const (
	// DirectionLTR merepresentasikan penulisan dari kiri ke kanan (Left-to-Right).
	DirectionLTR Direction = "ltr"

	// DirectionRTL merepresentasikan penulisan dari kanan ke kiri (Right-to-Left).
	DirectionRTL Direction = "rtl"
)

const (
	// LocaleAR merepresentasikan bahasa Arab (Arabic).
	LocaleAR Locale = "ar"
	// LocaleAZ merepresentasikan bahasa Azerbaijan.
	LocaleAZ Locale = "az"
	// LocaleBG merepresentasikan bahasa Bulgaria.
	LocaleBG Locale = "bg"
	// LocaleBN merepresentasikan bahasa Bengali.
	LocaleBN Locale = "bn"
	// LocaleCA merepresentasikan bahasa Katalan (Catalan).
	LocaleCA Locale = "ca"
	// LocaleCS merepresentasikan bahasa Ceko (Czech).
	LocaleCS Locale = "cs"
	// LocaleDA merepresentasikan bahasa Denmark.
	LocaleDA Locale = "da"
	// LocaleDE merepresentasikan bahasa Jerman (German).
	LocaleDE Locale = "de"
	// LocaleEL merepresentasikan bahasa Yunani (Greek).
	LocaleEL Locale = "el"
	// LocaleEN merepresentasikan bahasa Inggris (English).
	LocaleEN Locale = "en"
	// LocaleEO merepresentasikan bahasa Esperanto.
	LocaleEO Locale = "eo"
	// LocaleES merepresentasikan bahasa Spanyol (Spanish).
	LocaleES Locale = "es"
	// LocaleET merepresentasikan bahasa Estonia.
	LocaleET Locale = "et"
	// LocaleEU merepresentasikan bahasa Basque.
	LocaleEU Locale = "eu"
	// LocaleFA merepresentasikan bahasa Persia (Persian).
	LocaleFA Locale = "fa"
	// LocaleFI merepresentasikan bahasa Finlandia.
	LocaleFI Locale = "fi"
	// LocaleFR merepresentasikan bahasa Prancis (French).
	LocaleFR Locale = "fr"
	// LocaleGA merepresentasikan bahasa Irlandia (Irish).
	LocaleGA Locale = "ga"
	// LocaleGL merepresentasikan bahasa Galisia (Galician).
	LocaleGL Locale = "gl"
	// LocaleHE merepresentasikan bahasa Ibrani (Hebrew).
	LocaleHE Locale = "he"
	// LocaleHI merepresentasikan bahasa Hindi.
	LocaleHI Locale = "hi"
	// LocaleHU merepresentasikan bahasa Hongaria (Hungarian).
	LocaleHU Locale = "hu"
	// LocaleID merepresentasikan bahasa Indonesia.
	LocaleID Locale = "id"
	// LocaleIT merepresentasikan bahasa Italia (Italian).
	LocaleIT Locale = "it"
	// LocaleJA merepresentasikan bahasa Jepang (Japanese).
	LocaleJA Locale = "ja"
	// LocaleKO merepresentasikan bahasa Korea (Korean).
	LocaleKO Locale = "ko"
	// LocaleKY merepresentasikan bahasa Kirgiz (Kyrgyz).
	LocaleKY Locale = "ky"
	// LocaleLT merepresentasikan bahasa Lituania.
	LocaleLT Locale = "lt"
	// LocaleLV merepresentasikan bahasa Latvia.
	LocaleLV Locale = "lv"
	// LocaleMS merepresentasikan bahasa Melayu (Malay).
	LocaleMS Locale = "ms"
	// LocaleNB merepresentasikan bahasa Norwegia Bokmål.
	LocaleNB Locale = "nb"
	// LocaleNL merepresentasikan bahasa Belanda (Dutch).
	LocaleNL Locale = "nl"
	// LocalePTBR merepresentasikan bahasa Portugis (Brasil).
	LocalePTBR Locale = "pt-BR"
	// LocalePL merepresentasikan bahasa Polandia (Polish).
	LocalePL Locale = "pl"
	// LocalePT merepresentasikan bahasa Portugis (Portugal).
	LocalePT Locale = "pt"
	// LocaleRO merepresentasikan bahasa Rumania.
	LocaleRO Locale = "ro"
	// LocaleRU merepresentasikan bahasa Rusia.
	LocaleRU Locale = "ru"
	// LocaleSK merepresentasikan bahasa Slowakia.
	LocaleSK Locale = "sk"
	// LocaleSL merepresentasikan bahasa Slovenia.
	LocaleSL Locale = "sl"
	// LocaleSQ merepresentasikan bahasa Albania.
	LocaleSQ Locale = "sq"
	// LocaleSV merepresentasikan bahasa Swedia.
	LocaleSV Locale = "sv"
	// LocaleTH merepresentasikan bahasa Thailand.
	LocaleTH Locale = "th"
	// LocaleTL merepresentasikan bahasa Tagalog (Filipina).
	LocaleTL Locale = "tl"
	// LocaleTR merepresentasikan bahasa Turki.
	LocaleTR Locale = "tr"
	// LocaleUK merepresentasikan bahasa Ukraina.
	LocaleUK Locale = "uk"
	// LocaleUR merepresentasikan bahasa Urdu.
	LocaleUR Locale = "ur"
	// LocaleVI merepresentasikan bahasa Vietnam.
	LocaleVI Locale = "vi"
	// LocaleZHHans merepresentasikan bahasa Mandarin (Aksara Sederhana).
	LocaleZHHans Locale = "zh-Hans"
	// LocaleZHHant merepresentasikan bahasa Mandarin (Aksara Tradisional).
	LocaleZHHant Locale = "zh-Hant"
)

var Locales []Locale = []Locale{
	LocaleAR,
	LocaleAZ,
	LocaleBG,
	LocaleCA,
	LocaleCS,
	LocaleDA,
	LocaleDE,
	LocaleEL,
	LocaleEN,
	LocaleEO,
	LocaleES,
	LocaleET,
	LocaleEU,
	LocaleFA,
	LocaleFI,
	LocaleFR,
	LocaleGA,
	LocaleGL,
	LocaleHE,
	LocaleHI,
	LocaleHU,
	LocaleID,
	LocaleIT,
	LocaleJA,
	LocaleKO,
	LocaleKY,
	LocaleLT,
	LocaleLV,
	LocaleMS,
	LocaleNB,
	LocaleNL,
	LocalePL,
	LocalePT,
	LocalePTBR,
	LocaleRO,
	LocaleRU,
	LocaleSK,
	LocaleSL,
	LocaleSQ,
	LocaleSV,
	LocaleTH,
	LocaleTL,
	LocaleTR,
	LocaleUK,
	LocaleUR,
	LocaleVI,
	LocaleZHHans,
	LocaleZHHant,
}
