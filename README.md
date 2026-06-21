# jasakode/i18n

[![Go Reference](https://pkg.go.dev/badge/github.com/jasakode/i18n.svg)](https://pkg.go.dev/github.com/jasakode/i18n)
[![Go Report Card](https://goreportcard.com/badge/github.com/jasakode/i18n)](https://goreportcard.com/report/github.com/jasakode/i18n)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

`jasakode/i18n` adalah pustaka internasionalisasi (i18n) untuk Go (Golang) yang terinspirasi langsung dari keandalan dan kenyamanan developer (*developer experience*) yang ada pada **next-intl** di ekosistem Next.js. 

Didesain khusus untuk aplikasi modern yang membutuhkan performa tinggi, pustaka ini memanfaatkan fitur `embed` bawaan Go dan dioptimalkan agar berjalan sangat cepat dengan konsumsi memori (*memory allocation*) seminimal mungkin (sangat cocok disandingkan dengan **Fiber / Fasthttp**).

## ✨ Fitur Utama

- **🚀 Developer Experience Premium:** Inisialisasi super ringkas, langsung mengembalikan fungsi penerjemah tunggal (`t`).
- **📦 Zero External Dependencies:** Murni menggunakan pustaka standar Go dan fitur native `embed.FS`.
- **⚡ High Performance:** Optimisasi pencarian kata (*lookup dictionary*) yang dirancang untuk menangani jutaan *request* per detik tanpa membebani *Garbage Collector*.
- **🔧 Flexibel / Framework Agnostic:** Dapat digunakan langsung pada pustaka standar `net/http`, maupun framework populer seperti Fiber, Gin, Echo, dan lainnya.
- **🎯 Dynamic Interpolation (Upcoming):** Dukungan penyuntikan variabel dinamis ke dalam string bahasa (Contoh: `t("welcome", i18n.Map{"name": "Antonius"})`).

## ⚙️ Instalasi

Unduh pustaka ini ke proyek Go Anda menggunakan perintah:

```bash
go get github.com/jasakode/i18n
```


## Dictionary
```json
{
    "app": {
        "welcome": "Hello {Name}",
    }
}
```

```go
type People struct {
    Name string
}
john := People{ Name: "John" }
locale := i18n.New()
locale.T("app.welcome", i18n.LocaleAR, john) // John مرحبًا
```