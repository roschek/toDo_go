package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func newRouter() *mux.Router {

	r := mux.NewRouter()
	r.HandleFunc("/", handler).Methods("GET")
	r.HandleFunc("/note", getNoteHandler).Methods("GET")
	r.HandleFunc("/note", createNoteHandler).Methods("POST")
	staticFileDir := http.Dir("./static/")
	staticFileHandler := http.StripPrefix("/static/", http.FileServer(staticFileDir))
	r.PathPrefix("/static/").Handler(staticFileHandler).Methods("GET")
	return r
}

func main() {
	r := newRouter()
	http.ListenAndServe(":8080", r)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "We are starting at port :8080")
}
