package main

import "fmt"

type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("base with num=%v", b.num)
}

type container struct {
	base
	str string
}

type featuresMap map[string]any

type container2 struct {
	base
	featuresMap
}

func main() {

	// container tests

	co := container{
		base: base{
			num: 1,
		},
		str: "some name",
	}

	fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str)

	fmt.Println("also num:", co.base.num)

	fmt.Println("describe:", co.describe())

	type describer interface {
		describe() string
	}

	var d describer = co
	fmt.Println("describer:", d.describe())

	// container2 tests
	co2 := container2{
		base: base{
			num: 1,
		},
	}
	fmt.Println("Before features assignment: co2=", co2)
	features := featuresMap{
		"feature1": 0.001,
		"feature2": []float64{
			0.002, 0.003,
		},
	}

	co2.featuresMap = features

	fmt.Println("After features assignment: co2=", co2)
}
