package main

import (
	"fmt"               // digunakan untuk format I/O
	"io"                // untuk operasi I/O
	"net/http"          // package HTTP standar
	"net/http/httptest" // untuk membuat HTTP test server dan request
	"testing"           // package untuk unit test

	"github.com/julienschmidt/httprouter" // package router HTTP
	"github.com/stretchr/testify/assert"  // package assertion untuk testing
)

// TestParams adalah unit test untuk memastikan parameter path pada router dapat diambil dengan benar
func TestParams(t *testing.T) {

	// Membuat instance router baru
	router := httprouter.New()

	// Menambahkan route GET dengan parameter :id
	router.GET("/products/:id", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		// Mengambil parameter id dari path
		id := params.ByName("id")
		// Membuat response text dengan id produk
		text := "Product " + id
		// Menuliskan response ke writer
		fmt.Fprint(writer, text)
	})

	// Membuat request HTTP GET ke endpoint /products/1
	request := httptest.NewRequest("GET", "http://localhost:3000/products/1", nil)
	// Membuat response recorder untuk menangkap response
	recorder := httptest.NewRecorder()

	// Menjalankan router dengan request dan menangkap response
	router.ServeHTTP(recorder, request)

	// Mengambil hasil response
	response := recorder.Result()
	// Membaca body response
	body, _ := io.ReadAll(response.Body)

	// Melakukan assertion bahwa response sesuai dengan yang diharapkan
	assert.Equal(t, "Product 1", string(body))

}
