// Package main adalah package utama untuk testing HTTP Router
package main

// Import package-package yang dibutuhkan untuk testing
import (
	"fmt"
	"io"                // Package untuk operasi I/O
	"net/http"          // Package standar untuk HTTP
	"net/http/httptest" // Package untuk membuat HTTP test
	"testing"           // Package standar untuk unit testing

	"github.com/julienschmidt/httprouter" // Package untuk HTTP routing
	"github.com/stretchr/testify/assert"  // Package untuk assertion dalam testing
)

// TestMethodNotAllowed adalah fungsi test untuk menguji penanganan HTTP method yang tidak diizinkan
func TestMethodNotAllowed(t *testing.T) {
	// Inisialisasi router baru
	router := httprouter.New()

	// Mendefinisikan handler untuk method yang tidak diizinkan
	router.MethodNotAllowed = http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Gak Boleh") // Mengirim response "Gak Boleh" ketika method tidak diizinkan
	})

	// Mendefinisikan route untuk method POST
	router.POST("/", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		fmt.Fprint(writer, "POST") // Mengirim response "POST" ketika method POST dipanggil
	})

	// Membuat request testing dengan method PUT (yang tidak didefinisikan)
	request := httptest.NewRequest("PUT", "http://localhost:3000/", nil)
	// Membuat response recorder untuk menangkap response
	recorder := httptest.NewRecorder()

	// Menjalankan request melalui router
	router.ServeHTTP(recorder, request)

	// Mengambil hasil response
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body) // Membaca body response

	// Memastikan response body sesuai dengan yang diharapkan
	assert.Equal(t, "Gak Boleh", string(body))
}
