package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	// Wait some time
	time.Sleep(1 * time.Second)

	end := time.Since(start)
	clockTimeMilliseconds := int(end) / 1000 / 1000
	endMilliseconds := end.Milliseconds()

	fmt.Printf("clockTimeMilliseconds=%v\n", clockTimeMilliseconds)
	fmt.Printf("endMilliseconds      =%v\n", endMilliseconds)
}
