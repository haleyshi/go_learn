package main

import (
	"fmt"
	"log"
)

func main() {
	primitiveTypeSliceExec()
}

func primitiveTypeSliceExec() {
	fmt.Println("slice start")
	slice := NewSlice()
	slice.Add(1)
	fmt.Println("current slice", slice)
	slice.Add("hello")
	fmt.Println("current slice", slice)
	slice.Add(3)
	fmt.Println("current slice", slice)

	if err := slice.Add(2); err != nil {
		log.Print(err)
	}
	fmt.Println("current slice", slice)

	slice.Remove(3)
	fmt.Println("current slice", slice)
	slice.Remove(1)
	fmt.Println("current slice", slice)
	slice.Remove("hello")
	fmt.Println("current slice", slice)

	if err := slice.Remove("abc"); err != nil {
		log.Print(err)
	}
	fmt.Println("current slice", slice)
}
