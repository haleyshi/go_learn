package main

import "fmt"

func main() {
	nums := []int16{2, 3, 4}

	sum := int16(0)
	for _, num := range nums { // index, value: _ here means ignore the index
		sum += num
	}

	fmt.Println("sum:", sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Printf("nums[%d]=%d\n", i, num)
		}
	}

	kvs := map[string]string{"a": "apple", "b": "banana", "o": "orange"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	for i, c := range "go" { // for Unicode string, value is the ASCII code
		fmt.Println(i, c)
	}
}
