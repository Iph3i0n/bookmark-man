package main

import (
	"bookmark-man/models"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"bookmark-man/service"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()
	s := r.Host("localhost").Subrouter() //config
	s.HandleFunc("/", HomeHandler)
	s.HandleFunc("/bookmark", AddHandler).Methods("POST")                               //POST add
	s.HandleFunc("/bookmarks/{user}", UserBookmarksHandler)                             //GET all
	s.HandleFunc("/bookmarks/{user}/{id:[0-9]+}", EditBookmarksHandler).Methods("POST") //POST edit
	s.HandleFunc("/bookmarks/{user}/{id:[0-9]+}", BookmarksHandler).Methods("GET")
	http.Handle("/", r)

	log.Fatal(http.ListenAndServe("localhost:8080", r))
}

//home view all
//home view user bookmarks
//add / delete

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	io.WriteString(w, `{"alive": true}`)
}

func AddHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var b models.Bookmark
	err := decoder.Decode(&b)
	if err != nil {
		panic(err)
	}

	s := service.New()

	if err := s.Add("sophie", b); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "%v", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	io.WriteString(w, `{"success": true}`)
}

func EditBookmarksHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var b models.Bookmark
	err := decoder.Decode(&b)
	if err != nil {
		panic(err)
	}

	s := service.New()

	if _, err := s.Update("", b); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "%v", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	io.WriteString(w, `{"success": true}`)
}

func UserBookmarksHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	s := service.New()
	bm, err := s.GetBookmarksForUser(vars["user"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "%v", err)
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%v", bm)
}

func BookmarksHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	s := service.New()
	bm, err := s.Get(vars["user"], vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "%v", err)
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%v", bm)
}
