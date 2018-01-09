package main

import (
	"log"
	"net/http"
)

func main() {
	PORT := "127.0.0.1:8060"

	log.Print("Running server on " + PORT)

	http.HandleFunc("/", simpleFunc)
	http.HandleFunc("/static/", staticFunc)

	log.Fatal(http.ListenAndServe(PORT, nil))
}

func simpleFunc(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}

func staticFunc()
