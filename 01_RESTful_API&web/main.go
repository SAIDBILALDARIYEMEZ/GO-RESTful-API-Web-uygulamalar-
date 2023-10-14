package main

import (
	"fmt"
	"net/http"
)

// human struct, insan verilerini temsil eder
type human struct {
	FName string
	LName string
	Age   int
}

// ServeHTTP, HTTP isteklerini işleyen fonksiyon
func (h human) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// İnsan verilerini doldur
	h.FName = "Said"
	h.LName = "Darıyemez"
	h.Age = 20

	// Formu parse etmek için
	r.ParseForm()

	// Form bilgilerini görüntülemek için
	fmt.Println("Form Bilgisi:", r.Form)

	// URL'nin Path bilgisini almak için
	fmt.Println("Path:", r.URL.Path[1:])

	// Web tarayıcısına yanıt olarak bir HTML tablo gönder
	fmt.Fprint(w, "<table><tr><td><b>Ad</b></td><td><b>Soyad</b></td><td><b>Yaş</b></td></tr>"+
		"<tr><td>"+h.FName+"</td><td>"+h.LName+"</td><td>", h.Age, "</td></tr>"+
		"<tr></tr><tr></tr><tr><td><b>Path</b></td><td>"+r.URL.Path+"</td></tr></table>")
}

func main() {
	var h human
	err := http.ListenAndServe(":8080", h)
	CheckError(err)
}

// CheckError, hata kontrolü yapar ve gerekirse programı durdurur
func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
