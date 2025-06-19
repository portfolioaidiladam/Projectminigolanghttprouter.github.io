// router_pattern_test.go
// Contoh unit test untuk pola route pada httprouter
package main

import (
	"fmt"               // Untuk format string output
	"io"                // Untuk membaca response body
	"net/http"          // Package HTTP standar
	"net/http/httptest" // Untuk membuat HTTP request dan response palsu
	"testing"           // Package untuk unit test

	"github.com/julienschmidt/httprouter" // Library router HTTP
	"github.com/stretchr/testify/assert"  // Library assertion untuk testing
)

// TestRouterPatternNamedParameter menguji route dengan parameter bernama
func TestRouterPatternNamedParameter(t *testing.T) {

	// Membuat instance router baru
	router := httprouter.New()
	// Mendefinisikan route dengan parameter :id dan :itemId
	router.GET("/products/:id/items/:itemId", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		// Mengambil parameter id dari URL
		id := params.ByName("id")
		// Mengambil parameter itemId dari URL
		itemId := params.ByName("itemId")
		// Membuat response text
		text := "Product " + id + " Item " + itemId
		// Menulis response ke writer
		fmt.Fprint(writer, text)
	})

	// Membuat request palsu ke endpoint yang sudah didefinisikan
	request := httptest.NewRequest("GET", "http://localhost:3000/products/1/items/1", nil)
	// Membuat response recorder untuk menangkap response
	recorder := httptest.NewRecorder()

	// Menjalankan router dengan request dan recorder
	router.ServeHTTP(recorder, request)

	// Mengambil hasil response
	response := recorder.Result()
	// Membaca seluruh body response
	body, _ := io.ReadAll(response.Body)

	// Memastikan hasil response sesuai yang diharapkan
	assert.Equal(t, "Product 1 Item 1", string(body))

}

// TestRouterPatternCatchAllParameter menguji route dengan parameter catch-all (wildcard)
func TestRouterPatternCatchAllParameter(t *testing.T) {

	// Membuat instance router baru
	router := httprouter.New()
	// Mendefinisikan route dengan parameter wildcard *image
	router.GET("/images/*image", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		// Mengambil parameter image dari URL (semua setelah /images/)
		image := params.ByName("image")
		// Membuat response text
		text := "Image : " + image
		// Menulis response ke writer
		fmt.Fprint(writer, text)
	})

	// Membuat request palsu ke endpoint wildcard
	request := httptest.NewRequest("GET", "http://localhost:3000/images/small/profile.png", nil)
	// Membuat response recorder
	recorder := httptest.NewRecorder()

	// Menjalankan router dengan request dan recorder
	router.ServeHTTP(recorder, request)

	// Mengambil hasil response
	response := recorder.Result()
	// Membaca seluruh body response
	body, _ := io.ReadAll(response.Body)

	// Memastikan hasil response sesuai yang diharapkan
	assert.Equal(t, "Image : /small/profile.png", string(body))

}
