package main

import (
	"fmt"
	"sync"
	"time"
)

func doWork(id int) {
	fmt.Printf("Worker %d starting\n", id)

	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func main() {

	const workersCount = 5
	var wg sync.WaitGroup

	for i := 1; i <= workersCount; i++ {
		wg.Add(1)

		// Avoid re-use of the same i value in each goroutine closure. See the https://golang.org/doc/faq#closures_and_goroutines for more details.
		i := i

		fmt.Println("Launching worker", i)
		go func() {
			defer wg.Done()
			doWork(i)
		}()
	}

	fmt.Println("Waiting for", workersCount, "workers...")
	wg.Wait()
	fmt.Println("Main is done.")
}
