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

// TestNotFound menguji fungsi custom handler untuk halaman 404 (Not Found)
// Test ini memverifikasi bahwa ketika mengakses URL yang tidak terdaftar,
// router akan menampilkan pesan custom "Gak Ketemu" alih-alih halaman 404 default
func TestNotFound(t *testing.T) {

	// Membuat instance router baru dari httprouter
	router := httprouter.New()
	
	// Mengatur custom handler untuk halaman 404 (Not Found)
	// Handler ini akan dipanggil ketika URL yang diakses tidak terdaftar di router
	router.NotFound = http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		// Menulis pesan "Gak Ketemu" ke response body
		fmt.Fprint(writer, "Gak Ketemu")
	})

	// Membuat HTTP request untuk testing
	// Menggunakan URL yang tidak terdaftar di router untuk memicu 404 handler
	request := httptest.NewRequest("GET", "http://localhost:3000/404", nil)
	
	// Membuat response recorder untuk menangkap response dari handler
	recorder := httptest.NewRecorder()

	// Memanggil router untuk memproses request dan menghasilkan response
	router.ServeHTTP(recorder, request)

	// Mengambil response result dari recorder
	response := recorder.Result()
	
	// Membaca body response untuk mendapatkan konten yang ditulis oleh handler
	body, _ := io.ReadAll(response.Body)

	// Melakukan assertion untuk memverifikasi bahwa response body berisi "Gak Ketemu"
	// Ini memastikan bahwa custom 404 handler berfungsi dengan benar
	assert.Equal(t, "Gak Ketemu", string(body))

}
