package main

import "fmt"

type thing struct {
	myArray []float64
}

func main() {

	thing1 := thing{}

	fmt.Printf("thing1=%+v\n", thing1)
	fmt.Printf("thing1.myArray is nil=%t\n", thing1.myArray == nil)
	fmt.Printf("thing1.myArray is zero-length=%t\n", len(thing1.myArray) == 0)
	fmt.Printf("thing1.myArray has content=%t\n", len(thing1.myArray) > 0)

	var a1 []float64
	fmt.Printf("a1 is nil: %t\n", a1 == nil)
	fmt.Printf("a1 is zero-length: %t\n", len(a1) == 0)
	fmt.Printf("a1 has content=%t\n", len(a1) > 0)

	a2 := make([]float64, 0)
	fmt.Printf("a2 is nil: %t\n", a2 == nil)
	fmt.Printf("a2 is zero-length: %t\n", len(a2) == 0)
	fmt.Printf("a2 has content=%t\n", len(a2) > 0)
}
