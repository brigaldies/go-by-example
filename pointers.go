package main

import "fmt"

func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	i = 2
	zeroptr(i) // Must pass a pointer. Go does not automatically pass the pointer to "i".
	fmt.Println("zeroptr:", i)

	fmt.Println("pointer:", &i)
}
