package main

import (
	"fmt"
	"net/http"
)

func StartServer() {
	fmt.Println("Starting additional HTTP server on :8081...")
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})
	http.ListenAndServe(":8081", nil)
}
