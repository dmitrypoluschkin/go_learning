package main

import (
	"fmt"
)

var x int = 10
var y int = x * 2
var i int = 5
var a float64 = 4.4

var t float64 = float64(x) + a

func main() {
	fmt.Println("y = ", y)
	fmt.Println(y == i)
	fmt.Println(y != i)
	fmt.Println(t)
}
