package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var fileName string

func main() {

	var PORT string

	if len(os.Args) <= 1 {
		log.Fatal("Correct syntax - ./main <fileName> [<IP Address:PORT>]")
		os.Exit(1)
	} else {
		fileName = os.Args[1]

		if len(os.Args) > 2 {
			PORT = os.Args[2]
		} else {
			PORT = ":8060"
		}
	}

	log.Println("Running server on " + PORT)
	log.Println("Requested file - " + fileName)

	http.HandleFunc("/", staticFunc)

	log.Fatal(http.ListenAndServe(PORT, nil))
}

func staticFunc(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {

		// if file is present

		if data, err := ioutil.ReadFile(fileName); err == nil {
			w.Write(data)
		} else {
			w.WriteHeader(404)
			w.Write([]byte("The requested file does not exist"))
		}
	}
}
