package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type Response struct {
	Result string `json:"Result"`
	Error  string `json:"Error"`
}

type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

type Server struct {
	port string
}

var books []Book
var borrowed []Book

func NewServer(port string) *Server {
	return &Server{port: port}
}

func (s *Server) Start() error {
	http.HandleFunc("/book", s.handleBook)
	return http.ListenAndServe(":"+s.port, nil)
}

func (s *Server) handleBook(w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	case http.MethodPost:
		handlePost(w, r)

	case http.MethodGet:
		handleGet(w, r)

	case http.MethodPut:
		handlePut(w, r)

	case http.MethodDelete:
		handleDelete(w, r)

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

/* ---------------- POST ---------------- */

func handlePost(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		errorJSON(w, "title or author cannot be empty")
		return
	}

	title := strings.ToLower(r.FormValue("title"))
	author := strings.ToLower(r.FormValue("author"))

	if title == "" || author == "" {
		errorJSON(w, "title or author cannot be empty")
		return
	}

	book := Book{Title: title, Author: author}

	if exists(books, book) {
		okJSON(w, "this book is already in the library")
		return
	}

	books = append(books, book)
	okJSON(w, fmt.Sprintf("added book %s by %s", title, author))
}


/* ---------------- GET ---------------- */

func handleGet(w http.ResponseWriter, r *http.Request) {
	title := strings.ToLower(r.URL.Query().Get("title"))
	author := strings.ToLower(r.URL.Query().Get("author"))

	if title == "" || author == "" {
		errorJSON(w, "title or author cannot be empty")
		return
	}

	book := Book{Title: title, Author: author}

	if !exists(books, book) {
		errorJSON(w, "this book does not exist")
		return
	}

	if exists(borrowed, book) {
		errorJSON(w, "this book is borrowed")
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(book)
}


/* ---------------- PUT ---------------- */

func handlePut(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Query().Get("title")
	author := r.URL.Query().Get("author")

	if title == "" || author == "" {
		errorJSON(w, "title or author cannot be empty")
		return
	}

	var body struct {
		Borrow *bool `json:"borrow"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil || body.Borrow == nil {
		errorJSON(w, "borrow value cannot be empty")
		return
	}

	book := Book{Title: title, Author: author}

	if !exists(books, book) {
		errorJSON(w, "this book does not exist")
		return
	}

	if *body.Borrow {
		// borrow
		if exists(borrowed, book) {
			errorJSON(w, "this book is already borrowed")
			return
		}
		borrowed = append(borrowed, book)
		okJSON(w, "you have borrowed this book successfully")
	} else {
		// return
		if !exists(borrowed, book) {
			errorJSON(w, "this book is already in the library")
			return
		}
		remove(&borrowed, book)
		okJSON(w, "thank you for returning this book")
	}
}

/* ---------------- DELETE ---------------- */

func handleDelete(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Query().Get("title")
	author := r.URL.Query().Get("author")

	if title == "" || author == "" {
		errorJSON(w, "title or author cannot be empty")
		return
	}

	book := Book{Title: title, Author: author}

	if !exists(books, book) {
		errorJSON(w, "this book does not exist")
		return
	}

	remove(&books, book)
	remove(&borrowed, book)
	okJSON(w, "successfully deleted")
}

/* ---------------- Helpers ---------------- */

func exists(list []Book, b Book) bool {
	for _, v := range list {
		if equal(v, b) {
			return true
		}
	}
	return false
}

func remove(list *[]Book, b Book) {
	for i, v := range *list {
		if equal(v, b) {
			*list = append((*list)[:i], (*list)[i+1:]...)
			return
		}
	}
}

func equal(a, b Book) bool {
	return strings.EqualFold(a.Title, b.Title) &&
		strings.EqualFold(a.Author, b.Author)
}

func okJSON(w http.ResponseWriter, msg string) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Response{Result: msg, Error: ""})
}

func errorJSON(w http.ResponseWriter, msg string) {
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(Response{Result: "", Error: msg})
}

/* ---------------- main ---------------- */

func main() {
	s := NewServer("4001")
	s.Start()
}
