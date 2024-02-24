package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request)  {
	// Check if the current request URL path matches exactly "/"
	// If it doesn't the return 404 not found
	if r.URL.Path!= "/" {
		http.NotFound(w, r)
		return
	}

	w.Write([]byte("Hello World"))
}

func snippetView(w http.ResponseWriter, r *http.Request)  {
	w.Write([]byte("Snippet View"))
}

func snippetCreate(w http.ResponseWriter, r *http.Request)  {
	// Check whether the request method is POST
	if r.Method!= http.MethodPost {
		w.Header().Set("Allow", "POST")
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Write([]byte("Snippet Create"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	log.Println("Starting server on port 4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}