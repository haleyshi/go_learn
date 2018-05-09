package main

import (
	"fmt"
	"reflect"
)

// User struct
type User struct {
	id   int
	name string
	age  int
	sex  string
}

var value interface{} = &User{1, "Tom", 20, "male"}

func main() {
	v := reflect.ValueOf(value)
	fmt.Println(v)
}
