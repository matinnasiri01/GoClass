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
			ErrorResponse(w,"Bad Request!")
			panic(err)
		}

		if b, _ := check(books, book); b {
			ResultResponse(w,"this book is already in the library")
			return
		}

		books = append(books, book)
		ResultResponse(w, fmt.Sprintf("added book %s by %s", book.Title, book.Author))

	}

}

func main() {
	s := NewServer("4001")
	s.Start()
}

func check(l []Book, b Book) (bool, int) {
	if len(l) == 0 {
		return false, -1
	}
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

func ResultResponse(w http.ResponseWriter, message string) {
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(Response{Result: message})
}
func ErrorResponse(w http.ResponseWriter, message string) {
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(Response{Error: message})
}
