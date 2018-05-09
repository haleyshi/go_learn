package main

import (
	"fmt"
	"log"
)

// Student is a test struct
type Student struct {
	id   string
	name string
}

func (stu Student) equals(obj interface{}) bool {
	if student, ok := obj.(Student); ok {
		return stu.id == student.id
	}

	return false
}

func main() {
	commonSliceExec()
}

func commonSliceExec() {
	fmt.Println("slice start")
	slice := NewSlice()
	slice.Add(1)
	fmt.Println("current slice", slice)
	slice.Add("hello")
	fmt.Println("current slice", slice)
	slice.Add(Student{"001", "Tom"})
	fmt.Println("current slice", slice)

	if err := slice.Add(Student{"001", "Jerry"}); err != nil {
		log.Print(err)
	}
	fmt.Println("current slice", slice)

	slice.Remove(1)
	fmt.Println("current slice", slice)
	slice.Remove("hello")
	fmt.Println("current slice", slice)

	if err := slice.Remove(Student{"002", "Tom"}); err != nil {
		log.Print(err)
	}
	fmt.Println("current slice", slice)

	slice.Remove(Student{"001", "Jerry"})
	fmt.Println("current slice", slice)
}
