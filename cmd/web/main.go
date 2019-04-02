package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/course", showCourse)
	mux.HandleFunc("/course/create", createCourse)

	// File Server for serving "./ui/static" directory and it's files.
	fs := http.FileServer(http.Dir("./ui/static/"))

	// Register file server to handle URL paths with "/static".
	// Strip the "/static prefix before the request reachest the file server
	mux.Handle("/static/", http.StripPrefix("/static", fs))

	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
