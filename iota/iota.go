package main

import "fmt"

const (
	x = 1 << iota
	y = 3 << iota
	z
	m
	n = 2 << iota
	o = iota
)

func main() {
	const (
		a = iota
		b
		c
		d = "abc"
		e
		f = 100
		g
		h = iota
		i
	)

	fmt.Println(a, b, c, d, e, f, g, h, i)
	
	fmt.Println("x=", x)
	fmt.Println("y=", y)
	fmt.Println("z=", z)
	fmt.Println("m=", m)
	fmt.Println("n=", n)
	fmt.Println("o=", o)
}