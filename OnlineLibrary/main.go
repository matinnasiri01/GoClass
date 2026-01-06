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
	Title  string `josn:"title"`
	Author string `josn:"author"`
}

var borrow []Book

var books []Book

type Server struct {
	port string
}

func NewServer(port string) *Server {
	return &Server{port: port}
}

func (s *Server) Start() error {
	http.HandleFunc("/book", s.handelBook)

	addr := ":" + s.port
	fmt.Printf("Server starting on port %s\n", s.port)
	return http.ListenAndServe(addr, nil)
}

func (s *Server) handelBook(w http.ResponseWriter, r *http.Request) {

	// post
	if r.Method == "POST" {
		var book Book
		err := json.NewDecoder(r.Body).Decode(&book)
		if err != nil {
			fmt.Println(err)
		}

		ti, au := strings.ToLower(book.Title), strings.ToLower(book.Author)
		for _, s := range books {
			if s.Title == ti && s.Author == au {
				fmt.Println("Sik!")
				return

			}
		}
		books = append(books, book)
		w.WriteHeader(http.StatusAccepted)
		fmt.Println("Add!")

	}

	if r.Method == "GET" {
		qu := r.URL.Query()

		ti, at := qu.Get("title"), qu.Get("author")

		if ti == "" || at == "" {
			ErrorResponse(w,"title or author cannot be empty")
			return
		}
		if len(books) == 0 {
			ErrorResponse(w,"title or author cannot be empty")
			return
		}

		if ch, index := check(books, Book{Title: ti, Author: at}); ch {
			json.NewEncoder(w).Encode(books[index])
		} else {
			ErrorResponse(w,"title or author cannot be empty")
		}

	}
	if r.Method == "DELETE" {
		qu := r.URL.Query()
		ti, at := qu.Get("title"), qu.Get("author")
		if isDatabaseEmpty() {

		}
		for _, b := range books {
			if b.Title == ti && b.Author == at {
				json.NewEncoder(w).Encode(b)
			}
		}

	}
}

func main() {
	s := NewServer("4001")
	s.Start()
}

func isDatabaseEmpty() bool {
	return len(books) == 0
}

func check(l []Book, b Book) (bool, int) {
	low := func(s string) string {
		return strings.ToLower(s)
	}
	for i, e := range l {
		if low(e.Title) == low(b.Title) && low(e.Author) == low(b.Author) {
			return true, i
		}
	}
	return false, -1
}

	w.WriteHeader(http.StatusNotFound)
}
