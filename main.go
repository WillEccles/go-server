package main

import (
	"log"
	"net/http"
	"fmt"
)

const (
	rootdir = "./public"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// serve index.html
		http.ServeFile(w, r, fmt.Sprintf("%v/%v", rootdir, r.URL.Path))
		if r.URL.Path == "/" {
			r.URL.Path = "/index.html"
		}
		log.Printf("Attempting to serve %v to %v", r.URL.Path, r.RemoteAddr)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
