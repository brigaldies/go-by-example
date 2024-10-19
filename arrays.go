package main

import "fmt"

type Product struct {
	PartNumber int
}

func main() {

	var a [5]int
	fmt.Println("emp:", a)

	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	fmt.Println("len:", len(a))

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("b:", b)

	// array slices
	c := b[1:3]
	fmt.Println("c:", c)
	c[0] = 6
	c[1] = 7
	fmt.Println("b (after slice edits):", b)

	parts := [3]Product{
		{
			PartNumber: 1,
		},
		{
			PartNumber: 2,
		},
		{
			PartNumber: 3,
		},
	}

	fmt.Println("parts:", parts)

	partsSlice := parts[1:2]
	fmt.Println("partsSlice:", partsSlice)
	partsSlice[0].PartNumber = 99

	fmt.Println("parts (after slice edits:", parts)

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
