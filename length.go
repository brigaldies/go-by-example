package main

import "fmt"

func main() {
	var x []float64
	x = nil
	fmt.Printf("x=%v\n", x)
	fmt.Printf("len(x)=%d\n", len(x))
}
