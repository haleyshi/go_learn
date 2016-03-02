package main 

var a = "SGH"
var b string = "sgh"
var c bool

var x, y int
var (
	d int
	e bool
)

var f, g int = 1, 2
var h, i = 100, "hello"


func main() {
	j, k := 200, "world"  // This format can only exist in the body of func
	println(a, b, c, x, y, d, e, f, g, h, j, i, k)	
}