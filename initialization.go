package main

import "fmt"

var HealthProductsExcludedStoreLocations map[int]bool
var HealthProductsExcludeCatetories []int

func main() {
	fmt.Println("HealthProductsExcludedStoreLocations", HealthProductsExcludedStoreLocations)
	fmt.Println("HealthProductsExcludeCatetories", HealthProductsExcludeCatetories)
}
