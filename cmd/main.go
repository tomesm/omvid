package main

import (
	"log"
	"net/http"
)

// Define a home handler function which writes a byte slice containing response body
func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from OMVID"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)

}
