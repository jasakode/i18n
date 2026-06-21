// Package i18n provides a lightweight, high-performance internationalization (i18n)
// solution for Go applications, inspired by the elegant developer experience of next-intl.
// It leverages Go's native embed.FS to bundle translation files into a single binary.
package i18n

import (
	"database/sql/driver"
	"fmt"
)

// Locale defines a string type for internationalization (i18n) settings.
// It adheres to the IETF BCP 47 language tags standard, which builds upon
// the ISO 639-1 alpha-2 language codes and handles script/region variations
// (e.g., "zh-Hans" and "zh-Hant").
//
// See: https://tools.ietf.org/html/bcp47 (IETF BCP 47)
// See: https://www.iso.org/standard/74575.html (ISO 639-1)
//
// For a complete list of reference language codes, see:
// https://www.loc.gov/standards/iso639-2/php/code_list.php (Official Library of Congress List)
// https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes (Wikipedia Table)
type Locale string

// Direction mendefinisikan tipe data string untuk arah penulisan teks (Text Directionality).
type Direction string

// Language merepresentasikan informasi mendetail mengenai sebuah bahasa,
// termasuk identitas lokalisasi, arah teks, simbol bendera, dan penamaan.
type Language struct {
	// ID merupakan kode unik locale (misal: "id", "en", "zh-Hans").
	ID Locale `json:"id" xml:"id"`

	// Direction menentukan arah penulisan teks, apakah "ltr" atau "rtl".
	// Sangat penting untuk kebutuhan penyesuaian tata letak (layout layouting) di frontend.
	Direction Direction `json:"direction" xml:"direction"`

	// Flag menyimpan emoji bendera yang merepresentasikan bahasa/negara tersebut (misal: "🇮🇩", "🇬🇧").
	Flag string `json:"flag" xml:"flag"`

	// EnglishName adalah nama bahasa dalam format internasional / bahasa Inggris (misal: "Indonesian", "Arabic").
	EnglishName string `json:"english_name" xml:"english_name"`

	// NativeName adalah nama bahasa dalam penulisan asli bahasa tersebut (misal: "Bahasa Indonesia", "العربية").
	NativeName string `json:"native_name" xml:"native_name"`
}

// TimeZone mendefinisikan tipe data string untuk lokasi zona waktu berdasarkan IANA Time Zone Database.
// Contoh: "Asia/Jakarta", "America/New_York", "Europe/London".
type TimeZone string

// Currency mendefinisikan tipe data string untuk kode mata uang berdasarkan standar ISO 4217.
// Contoh: "IDR", "USD", "EUR", "JPY".
type Currency struct {
	Code string
	Name map[Locale]string
}

// return string
func (l Locale) String() string { return string(l) }

// Parse mencoba mengubah string input acak menjadi tipe Locale yang valid.
// Jika tidak ada kode yang cocok atau input tidak dikenali, ia akan langsung
// mengembalikan nilai defaultLocale yang ditentukan oleh pengguna.
func Parse(input string, defaultLocale Locale) Locale {
	if len(input) < 2 {
		return defaultLocale
	}

	// Logika Ekstraksi Universal (Mendukung string bersih maupun HTTP Header)
	mainLang := input
	for i := 0; i < len(input); i++ {
		if input[i] == ',' || input[i] == ';' {
			mainLang = input[:i]
			break
		}
	}

	if len(mainLang) > 0 && mainLang[0] == ' ' {
		mainLang = mainLang[1:]
	}

	if len(mainLang) < 2 {
		return defaultLocale
	}

	// Evaluasi pencocokan token secara universal
	switch mainLang {
	case "ar", "AR", "ar-EG", "ar-SA":
		return LocaleAR
	case "az", "AZ":
		return LocaleAZ
	case "bg", "BG":
		return LocaleBG
	case "bn", "BN":
		return LocaleBN
	case "ca", "CA":
		return LocaleCA
	case "cs", "CS", "cz", "CZ":
		return LocaleCS
	case "da", "DA":
		return LocaleDA
	case "de", "DE", "de-DE", "de-CH":
		return LocaleDE
	case "el", "EL":
		return LocaleEL
	case "en", "EN", "en-US", "en-GB", "en-AU":
		return LocaleEN
	case "eo", "EO":
		return LocaleEO
	case "es", "ES", "es-ES", "es-MX":
		return LocaleES
	case "et", "ET":
		return LocaleET
	case "eu", "EU":
		return LocaleEU
	case "fa", "FA":
		return LocaleFA
	case "fi", "FI":
		return LocaleFI
	case "fr", "FR", "fr-FR", "fr-CA":
		return LocaleFR
	case "ga", "GA":
		return LocaleGA
	case "gl", "GL":
		return LocaleGL
	case "he", "HE", "iw":
		return LocaleHE
	case "hi", "HI":
		return LocaleHI
	case "hu", "HU":
		return LocaleHU
	case "id", "ID", "id-ID", "in", "IN":
		return LocaleID
	case "it", "IT", "it-IT":
		return LocaleIT
	case "ja", "JA", "ja-JP":
		return LocaleJA
	case "ko", "KO", "ko-KR":
		return LocaleKO
	case "ky", "KY":
		return LocaleKY
	case "lt", "LT":
		return LocaleLT
	case "lv", "LV":
		return LocaleLV
	case "ms", "MS", "ms-MY":
		return LocaleMS
	case "nb", "NB", "no", "NO":
		return LocaleNB
	case "nl", "NL", "nl-NL":
		return LocaleNL
	case "pl", "PL":
		return LocalePL
	case "pt", "PT", "pt-PT":
		return LocalePT
	case "ro", "RO":
		return LocaleRO
	case "ru", "RU", "ru-RU":
		return LocaleRU
	case "sk", "SK":
		return LocaleSK
	case "sl", "SL":
		return LocaleSL
	case "sq", "SQ":
		return LocaleSQ
	case "sv", "SV", "sv-SE":
		return LocaleSV
	case "th", "TH", "th-TH":
		return LocaleTH
	case "tl", "TL", "fil", "FIL":
		return LocaleTL
	case "tr", "TR", "tr-TR":
		return LocaleTR
	case "uk", "UK", "uk-UA":
		return LocaleUK
	case "ur", "UR":
		return LocaleUR
	case "vi", "VI", "vi-VN":
		return LocaleVI
	}

	// Pengecekan variasi kompleks (Chinese & Brasil)
	if len(mainLang) >= 7 && (mainLang[:7] == "zh-Hans" || mainLang[:7] == "zh-hans" || mainLang[:7] == "ZH-HANS") {
		return LocaleZHHans
	}
	if len(mainLang) >= 7 && (mainLang[:7] == "zh-Hant" || mainLang[:7] == "zh-hant" || mainLang[:7] == "ZH-HANT") {
		return LocaleZHHant
	}
	if len(mainLang) >= 2 && (mainLang[:2] == "zh" || mainLang[:2] == "ZH") {
		if len(mainLang) >= 5 && (mainLang[3:] == "TW" || mainLang[3:] == "HK" || mainLang[3:] == "MO") {
			return LocaleZHHant
		}
		return LocaleZHHans
	}

	if len(mainLang) >= 5 && (mainLang == "pt-BR" || mainLang == "pt-br" || mainLang == "PT-BR") {
		return LocalePTBR
	}

	return defaultLocale
}

// Parse memproses string input acak dan langsung memperbarui nilai objek Locale saat ini.
// Jika proses parsing berhasil, ia mengembalikan true dan nilai objek diperbarui.
// Jika gagal, ia mengembalikan false dan nilai objek lama dipertahankan (sebagai default alami).
func (l *Locale) Parse(input string) bool {
	// Panggil fungsi murni dengan mengirimkan nilai dirinya sendiri (*l) sebagai fallback default!
	res := Parse(input, *l)

	// Jika hasilnya sama dengan nilai saat ini, ada 2 kemungkinan:
	// 1. Memang match dengan bahasa yang sama.
	// 2. Gagal match sehingga fungsi Parse mengembalikan nilai (*l) yang kita kirim tadi.
	// Kita periksa ulang ke fungsi Parse menggunakan sentinel value khusus untuk memastikan status kesuksesan.
	const sentinel = Locale("⚡_NOT_FOUND_⚡")
	if Parse(input, sentinel) == sentinel {
		return false // Gagal parse, nilai objek asli (*l) tidak berubah
	}

	*l = res
	return true // Sukses parse dan nilai berhasil diperbarui
}

// ==================================================================================================
// utility for locale
// ==================================================================================================

// Value mengimplementasikan interface driver.Valuer untuk mengonversi tipe Locale
// menjadi string yang dapat disimpan ke dalam database oleh GORM.
func (l Locale) Value() (driver.Value, error) {
	// Mengembalikan nilai string murni ke driver database
	return string(l), nil
}

// Scan mengimplementasikan interface sql.Scanner untuk membaca data dari database
// dan memasukkannya kembali ke dalam tipe kustom Locale.
func (l *Locale) Scan(value any) error {
	if value == nil {
		*l = ""
		return nil
	}

	// Memastikan data yang dibaca dari database bisa dikonversi ke string atau slice byte
	switch v := value.(type) {
	case string:
		*l = Locale(v)
		return nil
	case []byte:
		*l = Locale(string(v))
		return nil
	default:
		return fmt.Errorf("gagal memindai tipe %T ke dalam Locale", value)
	}
}

// ==================================================================================================
