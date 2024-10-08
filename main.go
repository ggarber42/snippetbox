package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Write([]byte("Hello app"))
}

func snippetView(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Display specific snippet..."))
}

func snippetCreate(w http.ResponseWriter, r *http.Request){

	if r.Method != "POST"{
		w.Header().Set("Allow", "POST")
		w.WriteHeader(405)
		w.Write([]byte("method not allowed"))
		return	
	}

	w.Write([]byte("Create new snippet..."))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	log.Printf("Start server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}

