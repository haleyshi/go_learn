package main

import (
	"fmt"
	"time"
)

// Test struct
type Test struct {
	name string
}

func main() {
	i := 2
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool, string:
			fmt.Printf("I'm a %T\n", t)
		case int, int16, int32, int64, uint, uint16, uint32, uint64:
			fmt.Printf("I'm an %T\n", t)
		default:
			fmt.Printf("Self Defined, value: %v, type: %T\n", i, t)
		}
	}

	whatAmI(true)
	whatAmI(1)
	whatAmI(int32(1))
	whatAmI("hello")
	whatAmI(make(chan int, 1))
	whatAmI(Test{"test"})
}
