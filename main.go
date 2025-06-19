package main

import (
	"fmt"      // Mengimpor package fmt untuk output format
	"net/http" // Mengimpor package net/http untuk server HTTP

	"github.com/julienschmidt/httprouter" // Mengimpor package httprouter untuk routing HTTP
)

func main() {
	// Membuat instance router baru dari httprouter
	router := httprouter.New()

	// Menambahkan route GET untuk path root ("/")
	// Handler ini akan menampilkan "Hello HttpRouter" ke client
	router.GET("/", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		fmt.Fprint(writer, "Hello HttpRouter") // Menulis response ke client
	})

	// Membuat konfigurasi server HTTP
	server := http.Server{
		Handler: router, // Menggunakan router sebagai handler
		Addr:    "localhost:3000", // Menentukan alamat dan port server
	}

	// Menjalankan server HTTP
	server.ListenAndServe()
}
