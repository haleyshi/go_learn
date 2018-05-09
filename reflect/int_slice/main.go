package main

import (
	"fmt"
	"log"
)

func main() {
	intSliceExec()
}

func intSliceExec() {
	fmt.Println("int slice start")
	slice := NewSlice()
	slice.Add(1)
	fmt.Println("current int slice", slice)
	slice.Add(2)
	fmt.Println("current int slice", slice)
	slice.Add(3)
	fmt.Println("current int slice", slice)

	if err := slice.Add(2); err != nil {
		log.Print(err)
	}
	fmt.Println("current int slice", slice)

	slice.Remove(3)
	fmt.Println("current int slice", slice)
	slice.Remove(1)
	fmt.Println("current int slice", slice)
	slice.Remove(2)
	fmt.Println("current int slice", slice)

	if err := slice.Remove(1); err != nil {
		log.Print(err)
	}
	fmt.Println("current int slice", slice)
}
