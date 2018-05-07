package main

import (
	"fmt"
)

func main() {
	var numbers = make([]int16, 3, 5) // length=3, capacity=5
	printSlice(numbers)

	var nilSlice []int16
	printSlice(nilSlice)

	array := []int16{0, 1, 2, 3, 4, 5, 6, 7, 8, 9} // length=capacity=10
	printSlice(array)

	slice1 := array[1:4] // index [1:4), length=3, capacity=9
	printSlice(slice1)

	slice2 := array[:3] // index [0:3), length=3, capacity=10
	printSlice(slice2)

	slice3 := array[4:] // index [4, len(array)=10), length=6, capacity=6
	printSlice(slice3)

	var slice4 []int16
	slice4 = append(nilSlice, 0) // nilSlice not changed, slice4 = [0]
	printSlice(nilSlice)
	printSlice(slice4)

	slice4 = append(slice4, 1) // slice4 = [0, 1]
	printSlice(slice4)

	slice4 = append(slice4, 2, 3, 4) // slice4 = [0, 1, 2, 3, 4]
	printSlice(slice4)

	slice5 := make([]int16, len(slice4), cap(slice4)*2) // length=5, capacity=cap(slice4) * 2
	copy(slice5, slice4)                                // cocpy(dest, src)
	printSlice(slice5)

	sliceTest()
	twoDimensionArray()
}

func printSlice(x []int16) {
	fmt.Printf("length=%d capacity=%d slice=%v\n", len(x), cap(x), x)
}

func sliceTest() {
	arr := []int16{1, 2, 3, 4, 5}
	s := arr[:]

	for e := range s {
		fmt.Println(s[e])
	}

	s1 := make([]int16, 3)
	for e := range s1 {
		fmt.Println(s1[e])
	}
}

func twoDimensionArray() {
	// 6 * 2
	var a = [][]int16{{0, 1}, {10, 11}, {20}, {30, 31}, {}, {50, 51}}
	var i, j int

	for i = 0; i < len(a); i++ {
		for j = 0; j < len(a[i]); j++ {
			fmt.Printf("a[%d][%d]=%d\n", i, j, a[i][j])
		}
	}
}
