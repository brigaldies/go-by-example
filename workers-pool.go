package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- string) {
	for j := range jobs {
		fmt.Println("doWork", id, "started  job", j)
		time.Sleep(time.Second)
		fmt.Println("doWork", id, "finished job", j)
		results <- fmt.Sprintf("Job %v: %v", j, j*2)
	}
}

func main() {

	const numWorkers = 3
	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan string, numJobs)

	for w := 1; w <= numWorkers; w++ {
		go worker(w, jobs, results)
	}

	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	for a := 1; a <= numJobs; a++ {
		result := <-results
		fmt.Println("Result:", result)
	}
}
