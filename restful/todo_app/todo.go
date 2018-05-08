package main

import "time"

// Todo struct
// struct tags used for json encoding
type Todo struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	Due       time.Time `json:"due"`
}

// Todos - Array of Todo
type Todos []Todo
