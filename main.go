package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handlerHTTP)

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatal(err)
	}
}

func handlerHTTP(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Hello there"))

	if err != nil {
		log.Fatal(err)
	}
}
