package main 

import "fmt"
import "unsafe"

const LENGTH int = 10
const WIDTH int = 5

const (
	Unknown = 0
	Female = 1
	Male = 2
)

const (
	x = "sgh"
	y = len(x)
	z = unsafe.Sizeof(x)
)

func main() {
	var area int
	const a, b, c = 1, true, "hello"

	area = LENGTH * WIDTH
	fmt.Printf("面积为：%d", area)
	println()
	println(a, b, c)
	println(x, y, z)
}