package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Index - Handle the HTTP Request /
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome")
}

// TodoIndex - Handle the HTTP Request /todos
func TodoIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(todos); err != nil {
		panic(err)
	}
}

// TodoShow - Handle the HTTP Request /todos/{todoId}
func TodoShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoID, ok := vars["todoId"]

	if ok {
		id, err := strconv.ParseInt(todoID, 10, 64)

		if err != nil {
			w.Header().Set("Content-Type", "application/text; charset=UTF-8")
			w.WriteHeader(http.StatusNotAcceptable)
			fmt.Fprintf(w, "The ID %s is not an integer.", todoID)
			return
		}

		t := RepoFindTodo(int(id))
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)

		if err := json.NewEncoder(w).Encode(t); err != nil {
			panic(err)
		}
	} else {
		fmt.Fprintln(w, "Should never happen.")
	}
}

// TodoCreate create a Todo from POST with JSON input
func TodoCreate(w http.ResponseWriter, r *http.Request) {
	var todo Todo
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1024*1024))

	if err != nil {
		panic(err)
	}

	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	if err := json.Unmarshal(body, &todo); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusUnprocessableEntity)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	t := RepoCreateTodo(todo)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(t); err != nil {
		panic(err)
	}
}

// TodoDelete - Handle the DELETE /todos/{todoId}
func TodoDelete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoID, ok := vars["todoId"]

	w.Header().Set("Content-Type", "application/text; charset=UTF-8")

	if ok {
		id, err := strconv.ParseInt(todoID, 10, 64)

		if err != nil {
			w.WriteHeader(http.StatusNotAcceptable)
			fmt.Fprintf(w, "The ID %s is not an integer.", todoID)
			return
		}

		if err := RepoDeleteTodo(int(id)); err != nil {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "The TODO with ID %s is not found.", todoID)
		} else {
			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, "The TODO with ID %s is deleted.", todoID)
		}
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Should never happen.")
	}
}
