// Package main berisi implementasi testing untuk middleware HTTP router
package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

// LogMiddleware adalah struct yang mengimplementasikan interface http.Handler
// untuk mencatat (logging) setiap request yang masuk
type LogMiddleware struct {
	// Handler adalah instance dari HTTP handler yang akan dijalankan setelah logging
	http.Handler
}

// ServeHTTP adalah method yang mengimplementasikan interface http.Handler
// Method ini akan mencatat pesan setiap kali ada request masuk dan
// kemudian meneruskan request tersebut ke handler berikutnya
func (middleware *LogMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Receive Request")
	middleware.Handler.ServeHTTP(writer, request)
}

// TestMiddleware adalah fungsi testing untuk memverifikasi bahwa middleware
// berfungsi dengan benar dalam melakukan logging dan meneruskan request
func TestMiddleware(t *testing.T) {
	// Inisialisasi router baru
	router := httprouter.New()
	
	// Mendaftarkan handler untuk path "/" dengan method GET
	router.GET("/", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		fmt.Fprint(writer, "Middleware")
	})

	// Membungkus router dengan LogMiddleware
	middleware := LogMiddleware{router}

	// Membuat request testing palsu ke "http://localhost:3000/"
	request := httptest.NewRequest("GET", "http://localhost:3000/", nil)
	
	// Membuat response recorder untuk menangkap response
	recorder := httptest.NewRecorder()

	// Menjalankan middleware dengan request palsu
	middleware.ServeHTTP(recorder, request)

	// Mengambil response yang dihasilkan
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	// Memverifikasi bahwa response body sesuai dengan yang diharapkan
	assert.Equal(t, "Middleware", string(body))
}
