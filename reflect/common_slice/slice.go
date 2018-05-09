package main

import (
	"fmt"
)

// Slice is type of int slice
type Slice []interface{}

// Comparable interface, all implementers should to implement equals method
type Comparable interface {
	equals(obj interface{}) bool
}

// NewSlice make an empty Slice
func NewSlice() Slice {
	return make(Slice, 0)
}

func equals(a, b interface{}) bool {
	if comparable, ok := a.(Comparable); ok {
		return comparable.equals(b)
	}

	return a == b
}

// Add add an element to Slice if it is not exist
func (slice *Slice) Add(element interface{}) error {
	for _, v := range *slice {
		if equals(v, element) {
			fmt.Printf("Slice: Add Element - %v already exist.\n", element)
			return ERR_ELEMENT_EXIST
		}
	}

	*slice = append(*slice, element)
	fmt.Printf("Slice: Add Element - %v added successfully.\n", element)
	return nil
}

// Remove remove an element from Slice
func (slice *Slice) Remove(element interface{}) error {
	found := false
	for i, v := range *slice {
		if equals(v, element) {
			*slice = append((*slice)[:i], (*slice)[i+1:]...)
			found = true
			break
		}
	}

	if !found {
		fmt.Printf("Slice: Remove Element - %v not exist.\n", element)
		return ERR_ELEMENT_NOT_EXIST
	}

	fmt.Printf("Slice: Remove Element - %v removed successfully.\n", element)
	return nil
}
