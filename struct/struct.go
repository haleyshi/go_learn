package main

import (
	"fmt"
)

type Book struct {
	title   string
	author  string
	subject string
	id      int
}

func main() {
	var Book1 Book
	var Book2 Book

	Book1.title = "Go"
	Book1.author = "Denis"
	Book1.subject = "Go"
	Book1.id = 1234567

	Book2.title = "Python"
	Book2.author = "James"
	Book2.subject = "Python"
	Book2.id = 9876543

	fmt.Println(Book1)
	fmt.Println(Book2)

	fmt.Printf("Author of Book1: %s\n", Book1.author)

	var bptr *Book
	bptr = &Book2
	bptr.author = "William"
	fmt.Printf("Author of Book2: %s\n", bptr.author)
}
