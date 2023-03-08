package main

import (
	"net/http"
	"log"
)

func echoHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("Error parsing form"))
		return
	}
	if r.Method != http.MethodGet {
		w.WriteHeader(405)
		w.Write([]byte("Invalid method"))
		return
	}
	f := r.Form
	if !f.Has("data") {
		w.WriteHeader(400)
		w.Write([]byte("Must specify data param"))
		return
	}
	w.WriteHeader(200)
	w.Write([]byte(f.Get("data")))
	return
}

func main() {
	http.HandleFunc("/", echoHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
