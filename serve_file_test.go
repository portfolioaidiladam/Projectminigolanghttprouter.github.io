// Package main mendefinisikan paket utama
package main

// Import package yang diperlukan
import (
	"embed"             // Untuk embed file statis
	"io"                // Untuk membaca response body
	"io/fs"             // Untuk manipulasi filesystem
	"net/http"          // Untuk HTTP server
	"net/http/httptest" // Untuk testing HTTP handler
	"testing"           // Untuk unit test

	"github.com/julienschmidt/httprouter" // Router eksternal
	"github.com/stretchr/testify/assert"  // Untuk assertion pada test
)

//go:embed resources
var resources embed.FS

// TestServerFile menguji serve file hello.txt
func TestServerFile(t *testing.T) {

	router := httprouter.New() // Membuat instance router baru
	directory, _ := fs.Sub(resources, "resources") // Mengambil subdirektori 'resources' dari embed FS
	router.ServeFiles("/files/*filepath", http.FS(directory)) // Routing untuk serve file statis

	request := httptest.NewRequest("GET", "http://localhost:3000/files/hello.txt", nil) // Membuat request GET ke file hello.txt
	recorder := httptest.NewRecorder() // Membuat response recorder untuk menangkap response

	router.ServeHTTP(recorder, request) // Menjalankan router dengan request yang dibuat

	response := recorder.Result() // Mengambil hasil response
	body, _ := io.ReadAll(response.Body) // Membaca seluruh body response

	assert.Equal(t, "Hello HttpRouter", string(body)) // Assert isi body sesuai ekspektasi

}

// TestServerFileGoodBye menguji serve file goodbye.txt
func TestServerFileGoodBye(t *testing.T) {

	router := httprouter.New() // Membuat instance router baru
	directory, _ := fs.Sub(resources, "resources") // Mengambil subdirektori 'resources' dari embed FS
	router.ServeFiles("/files/*filepath", http.FS(directory)) // Routing untuk serve file statis

	request := httptest.NewRequest("GET", "http://localhost:3000/files/goodbye.txt", nil) // Membuat request GET ke file goodbye.txt
	recorder := httptest.NewRecorder() // Membuat response recorder untuk menangkap response

	router.ServeHTTP(recorder, request) // Menjalankan router dengan request yang dibuat

	response := recorder.Result() // Mengambil hasil response
	body, _ := io.ReadAll(response.Body) // Membaca seluruh body response

	assert.Equal(t, "Good Bye HttpRouter", string(body)) // Assert isi body sesuai ekspektasi

}
