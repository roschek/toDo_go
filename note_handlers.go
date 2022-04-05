package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/lib/pq"
	"net/http"
)

type Note struct {
	Title string `json:"title"`
	Topic string `json:"topic"`
}

var notes []Note

func getNoteHandler(w http.ResponseWriter, r *http.Request) {
	notesListBytes, err := json.Marshal(notes)
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(notesListBytes)
}

func createNoteHandler(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("postgres", "user=theUser dbname=theDbName sslmode=verify-full")
	if err != nil {
		panic(err)
	}
	note := Note{}

	er := r.ParseForm()
	if er != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Println(r.Form)
	note.Title = r.Form.Get("title")
	note.Topic = r.Form.Get("topic")
	fmt.Println(r.Form)
	notes = append(notes, note)

	defer db.Close()
	http.Redirect(w, r, "/static/", http.StatusFound)
}
