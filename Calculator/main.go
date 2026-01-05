package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type Response struct {
	Result string `json:"result"`
	Error  string `json:"error"`
}

type Server struct {
	port string
}

func NewServer(port string) *Server {
	return &Server{port: port}
}

func (s *Server) Start() error {
	http.HandleFunc("/add", s.handleAdd)
	http.HandleFunc("/sub", s.handleSub)

	addr := ":" + s.port
	fmt.Printf("Server starting on port %s\n", s.port)
	return http.ListenAndServe(addr, nil)
}

func (s *Server) handleAdd(w http.ResponseWriter, r *http.Request) {
	s.handleCalculation(w, r, true)
}

func (s *Server) handleSub(w http.ResponseWriter, r *http.Request) {
	s.handleCalculation(w, r, false)
}

func (s *Server) handleCalculation(w http.ResponseWriter, r *http.Request, isAdd bool) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != "GET" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{Error: "Only GET method is allowed"})
		return
	}

	numbersParam := r.URL.Query().Get("numbers")
	if numbersParam == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{Error: "'numbers' parameter missing"})
		return
	}

	numbersStr := strings.Split(numbersParam, ",")
	if len(numbersStr) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{Error: "No numbers provided"})
		return
	}

	var result int64
	var hasFirstNumber bool

	for i, numStr := range numbersStr {
		num, err := strconv.ParseInt(strings.TrimSpace(numStr), 10, 64)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(Response{Error: fmt.Sprintf("Invalid number: %s", numStr)})
			return
		}

		if i == 0 {
			result = num
			hasFirstNumber = true
			continue
		}

		if isAdd {
			if num > 0 && result > 0 && result > 9223372036854775807-num {
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(Response{Error: "Overflow"})
				return
			}
			if num < 0 && result < 0 && result < -9223372036854775808-num {
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(Response{Error: "Overflow"})
				return
			}
			result += num
		} else {
			if num < 0 && result > 0 && result > 9223372036854775807+num {
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(Response{Error: "Overflow"})
				return
			}
			if num > 0 && result < 0 && result < -9223372036854775808+num {
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(Response{Error: "Overflow"})
				return
			}
			result -= num
		}
	}

	if !hasFirstNumber {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{Error: "No valid numbers provided"})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Response{
		Result: fmt.Sprintf("The result of your query is: %d", result),
		Error:  "",
	})
}


var testServer *Server
func getServer() *Server {
	if testServer == nil {
		testServer = NewServer("4001")
		go testServer.Start()
	}
	return testServer
}

func main() {
	getServer()
	resp, err := http.DefaultClient.Get("http://127.0.0.1:4001" + "/add?numbers=1,2")
	fmt.Print(resp)
	if  err != nil  {
		fmt.Errorf(err.Error())
	}
	defer resp.Body.Close()
}

