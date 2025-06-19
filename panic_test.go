// Package main berisi test untuk mengecek panic handler pada httprouter
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

// TestPanicHandler menguji fungsi panic handler pada httprouter
// Test ini memastikan bahwa ketika terjadi panic dalam handler,
// panic handler akan menangkap error dan mengembalikan response yang sesuai
func TestPanicHandler(t *testing.T) {

	// Membuat instance router baru dari httprouter
	router := httprouter.New()

	// Menentukan panic handler yang akan dipanggil ketika terjadi panic
	// Panic handler menerima 3 parameter: writer, request, dan error interface{}
	router.PanicHandler = func(writer http.ResponseWriter, request *http.Request, error interface{}) {
		// Menulis response ke writer dengan format "Panic : {error message}"
		fmt.Fprint(writer, "Panic : ", error)
	}

	// Mendaftarkan route GET "/" dengan handler yang sengaja memicu panic
	// Handler ini akan memanggil panic("Ups") setiap kali route ini diakses
	router.GET("/", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		// Sengaja memicu panic dengan pesan "Ups" untuk testing
		panic("Ups")
	})

	// Membuat HTTP request test untuk method GET ke endpoint "/"
	// httptest.NewRequest digunakan untuk membuat request yang dapat digunakan dalam testing
	request := httptest.NewRequest("GET", "http://localhost:3000/", nil)
	
	// Membuat HTTP response recorder untuk menangkap response dari handler
	// httptest.NewRecorder berfungsi sebagai http.ResponseWriter yang dapat dibaca
	recorder := httptest.NewRecorder()

	// Memanggil router untuk memproses request dan menulis response ke recorder
	// ServeHTTP akan memanggil handler yang sesuai dan panic handler jika terjadi panic
	router.ServeHTTP(recorder, request)

	// Mengambil hasil response dari recorder
	response := recorder.Result()
	
	// Membaca body response sebagai byte slice
	// io.ReadAll akan membaca semua data dari response body
	body, _ := io.ReadAll(response.Body)

	// Menggunakan testify/assert untuk memverifikasi bahwa response body
	// sama dengan string yang diharapkan: "Panic : Ups"
	// Jika tidak sama, test akan gagal
	assert.Equal(t, "Panic : Ups", string(body))

}
