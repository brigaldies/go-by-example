package main

import "fmt"

func main() {

	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:", m)

	v1 := m["k1"]
	fmt.Println("v1: ", v1)

	fmt.Println("len:", len(m))

	delete(m, "k2")
	fmt.Println("map:", m)

	k2Val, prs := m["k2"]
	fmt.Println("prs:", prs)
	if !prs {
		fmt.Println("Key 'k2' does not exist; k2Val:", k2Val)
	}

	n := map[string]string{"foo": "1", "bar": "2"}
	fmt.Println("map:", n)
	if val := n["baz"]; val == "" {
		fmt.Println("baz does not exist in map n")
	}
}
