package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// Todo struct
// struct tags used for json encoding
type Todo struct {
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	Due       time.Time `json:"due"`
}

// Todos - Array of Todo
type Todos []Todo

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/todos", TodoIndex)
	router.HandleFunc("/todos/{todoId}", TodoShow)
	log.Fatal(http.ListenAndServe(":8080", router))
}

// Index - Handle the HTTP Request /
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome")
}

// TodoIndex - Handle the HTTP Request /todos
func TodoIndex(w http.ResponseWriter, r *http.Request) {
	todos := Todos{
		Todo{Name: "Coding"},
		Todo{Name: "Testing"},
	}
	json.NewEncoder(w).Encode(todos)
}

// TodoShow - Handle the HTTP Request /todos/{todoId}
func TodoShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoID, ok := vars["todoId"]

	if ok {
		fmt.Fprintln(w, "Todo Show:", todoID)
	} else {
		fmt.Fprintln(w, "Should never happen.")
	}
}
