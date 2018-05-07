package main

import "fmt"

func main() {
	var a = 10
	fmt.Printf("Address of A: %x\n", &a)

	var ip *int
	ip = &a
	fmt.Printf("Address of ip point to: %x\n", ip)
	fmt.Printf("Value of ip point to: %d\n", *ip)

	var nilPtr *float32 // nil pointer
	fmt.Printf("Address of nil pointer: %x\n", nilPtr)

	if nilPtr == nil {
		fmt.Println("nilPtr is a nil pointer")
	}

	// pointer to pointer
	var ipp **int
	ipp = &ip
	fmt.Printf("ipp address: %x value: %x **ipp: %d\n", ipp, *ipp, **ipp)
}
