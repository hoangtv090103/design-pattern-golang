package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		if j < 0 {
			fmt.Println(j < 0)
		}
		fmt.Println("Worker", id, "started job", j, "...")
		time.Sleep(time.Second)
		fmt.Println("Worker", id, "finished job", j)
		results <- j * 2
	}
}

func main() {
	const numJobs = 5

	// Create a channel that allows us to send a unit of work to a worker pool
	jobs := make(chan int, numJobs) // Channel for job

	// Create a channel to send results to
	results := make(chan int, numJobs) // Channel for result

	// Spawn 3 workers
	for i := 1; i <= 3; i++ { // 3 workers
		go worker(i, jobs, results)
	}

	// Send units of work to the jobs channel
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}

	close(jobs)

	for a := 1; a <= numJobs; a++ {
		<-results
	}
}
