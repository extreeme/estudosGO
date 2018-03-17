package main

import (
	"fmt"
	m "math"
)

const (
	a = 1
	b = 2
)

func main() {
	var raio = 5.5
	const PI float64 = 3.1415
	area := PI * m.Pow(raio, 2)

	c, d, e, f := 1, false, 3, "opa!"

	fmt.Println(a, b)
	fmt.Println(c, d, e, f)
	fmt.Println("A área é: ", area)
}
