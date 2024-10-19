package main

import "fmt"

func main() {
	a := make([]int, 5)
	printSlice("a", a)

	b := make([]int, 0, 5)
	printSlice("b", b)

	// This crashes: "panic: runtime error: index out of range [0] with length 0"
	//b[0] = 0
	//b[1] = 0
	//b[2] = 0

	b = append(b, 0)
	b = append(b, 1)
	b = append(b, 2)

	printSlice("b (after assignments)", b)

	b = make([]int, 0, 5)
	printSlice("b", b)

	c := b[:2]
	printSlice("c", c)

	d := c[2:5]
	printSlice("d", d)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
