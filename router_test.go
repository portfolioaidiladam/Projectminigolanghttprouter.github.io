package main

import (
	"fmt"               // Package untuk format I/O
	"io"                // Package untuk operasi I/O
	"net/http"          // Package HTTP client dan server
	"net/http/httptest" // Package untuk testing HTTP
	"testing"           // Package untuk unit test

	"github.com/julienschmidt/httprouter" // Package router HTTP ringan
	"github.com/stretchr/testify/assert"  // Package assertion untuk testing
)

// TestRouter adalah unit test untuk memastikan router bekerja dengan benar
func TestRouter(t *testing.T) {

	router := httprouter.New() // Membuat instance router baru
	router.GET("/", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		fmt.Fprint(writer, "Hello World") // Mengirimkan response "Hello World" ke client
	})

	request := httptest.NewRequest("GET", "http://localhost:3000/", nil) // Membuat request HTTP GET palsu
	recorder := httptest.NewRecorder() // Membuat response recorder untuk menangkap output

	router.ServeHTTP(recorder, request) // Menjalankan router dengan request dan merekam response

	response := recorder.Result() // Mengambil hasil response
	body, _ := io.ReadAll(response.Body) // Membaca seluruh body response

	assert.Equal(t, "Hello World", string(body)) // Memastikan response body sesuai dengan yang diharapkan

}
