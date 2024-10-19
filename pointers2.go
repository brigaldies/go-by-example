package main

import "fmt"

type Widget struct {
	Name       string
	PartNumber int
	Price      float32
	Sizes      []float64
}

func main() {
	widget := Widget{}

	fmt.Printf("widget1: %+v\n", widget)
	fmt.Printf("widget1 name pointer: %p\n", &widget.Name)
	fmt.Printf("widget1 name value: %q\n", *(&widget.Name))

	fmt.Printf("widget1 part number pointer: %p\n", &widget.PartNumber)
	fmt.Printf("widget1 part number value: %d\n", *(&widget.PartNumber))

	pointer1 := &(widget.Sizes)
	fmt.Printf("widget1 sizes pointer: %p\n", pointer1)
	fmt.Printf("widget1 sizes pointer is nil: %t\n", pointer1 == nil)
	fmt.Printf("widget1 part number value: %v\n", *pointer1)
	*pointer1 = []float64{1.1, 2.2}
	fmt.Printf("widget: %+v\n", widget)

	widget2 := Widget{}
	pointer2 := &(widget2.Sizes)
	fmt.Printf("widget2 sizes pointer2: %p\n", pointer2)

	fmt.Printf("widgets' sizes pointers p1 == p2: %t\n", pointer1 == pointer2)
}
