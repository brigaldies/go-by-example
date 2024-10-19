package main

import (
	"fmt"
	"math"
)

type Hit struct {
	ProductId string `json:"product_id"`
}

func main() {

	hitsA := []Hit{
		{
			ProductId: "1000",
		},
		{
			ProductId: "1001",
		},
		{
			ProductId: "1002",
		},
	}

	hitsB := []Hit{
		{
			ProductId: "1002",
		},
		{
			ProductId: "1004",
		},
		{
			ProductId: "1000",
		},
	}

	fmt.Printf("NDCG@10=%v\n", getNDCG(10, hitsA, hitsB))
	fmt.Printf("NDCG@3=%v\n", getNDCG(3, hitsA, hitsB))
	fmt.Printf("NDCG@2=%v\n", getNDCG(2, hitsA, hitsB))

}

func getNDCG(total int, orderedHits, hitsToCheck []Hit) float64 {
	if len(orderedHits) == 1 && len(hitsToCheck) == 1 {
		return 1
	}
	if total > len(orderedHits) {
		total = len(orderedHits)
	}

	if total > len(hitsToCheck) {
		total = len(hitsToCheck)
	}

	fmt.Printf("total=%d\n", total)

	og := make(map[string]float64)
	var sumDCGMax float64
	for i := 0; i < total; i++ {
		v := float64(total - i)
		og[orderedHits[i].ProductId] = v
		gain := (math.Pow(2, v) - 1) / math.Log2(float64(2+i))
		fmt.Printf("iDCG Gain: %f\n", gain)
		sumDCGMax += gain
	}

	fmt.Printf("og=%v\n", og)
	fmt.Printf("iDCG=%f\n", sumDCGMax)

	var sumDCG float64
	for i := 0; i < total; i++ {
		v := og[hitsToCheck[i].ProductId]
		fmt.Printf("HitsB: i=%d, product id=%s ==> v=%f\n", i, hitsToCheck[i].ProductId, v)
		gain := (math.Pow(2, v) - 1) / math.Log2(float64(2+i))
		fmt.Printf("DCG Gain: %f\n", gain)
		sumDCG += gain
	}

	fmt.Printf("DCG=%f\n", sumDCG)

	nDCG := sumDCG / sumDCGMax

	fmt.Printf("og['key does not exist']=%v\n", og["key does not exist"])

	return nDCG
}
