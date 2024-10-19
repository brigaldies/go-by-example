package main

import (
	"fmt"
	"math/rand"
)

func main() {

	rand.Seed(10)

	rangeCenter := 500
	rangeSpan := 400

	randomRangeFrom := rangeCenter - rangeSpan/2
	if randomRangeFrom < 0 {
		randomRangeFrom = 0
	}
	randomRangeTo := rangeCenter + rangeSpan/2

	fmt.Printf("rangeCenter=%d, rangeSpan=%d, rangeFrom=%d, rangeTo=%d\n", rangeCenter, rangeSpan, randomRangeFrom, randomRangeTo)

	loopCount := 10
	for i := 0; i < loopCount; i++ {
		randowDraw := rand.Intn(randomRangeTo-randomRangeFrom+1) + randomRangeFrom
		fmt.Printf("[%d] Randow draw: %d\n", i, randowDraw)
	}
}
