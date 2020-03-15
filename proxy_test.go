package main

import (
	"log"
	"net/http"
	"testing"
)

func TestProxy(t *testing.T) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello 9091"))
	})
	log.Fatal(http.ListenAndServe(":9091", nil))
}
