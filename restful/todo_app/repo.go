package main

import "fmt"

var currentID int
var todos Todos

func init() {
	fmt.Println("initialize repo data...")
	RepoCreateTodo(Todo{Name: "Coding"})
	RepoCreateTodo(Todo{Name: "Testing"})
}

// RepoCreateTodo set the ID to Todo and append it to the Todo Array
func RepoCreateTodo(t Todo) Todo {
	currentID++
	t.ID = currentID
	todos = append(todos, t)
	return t
}

// RepoFindTodo get the Todo according to ID
func RepoFindTodo(id int) Todo {
	for _, t := range todos {
		if t.ID == id {
			return t
		}
	}

	// return empty Todo if not found
	return Todo{}
}

// RepoDeleteTodo delete a Todo from Array
func RepoDeleteTodo(id int) error {
	for i, t := range todos {
		if t.ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			return nil
		}
	}

	return fmt.Errorf("Could not find Todo with id of %d to delete", id)
}
