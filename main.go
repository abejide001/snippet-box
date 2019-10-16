package main

import (
	"log"
	"net/http"
)

func home(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Hello from here")) // returns "Hello from here as the response"
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	log.Println("Starting on port 4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
