package main

import (
	"fmt"
	"time"
)

func main() {
	numJobs := 10
	jobsChan := make(chan int, numJobs)
	completedJobsChan := make(chan int, numJobs)

	// 3 workers. Each worker will be a goroutine
	for w := 1; w <= 3; w++ {
		go worker(w, jobsChan, completedJobsChan)
	}

	// Load the jobsChan channel with job numbers
	for j := 1; j <= numJobs; j++ {
		jobsChan <- j
	}
	close(jobsChan)

	// Clear the channel and delay termination of the program until all jobs are finished
	for a := 1; a <= numJobs; a++ {
		<-completedJobsChan
	}
	close(completedJobsChan)

}

func worker(id int, jobsChan <-chan int, completedJobsChan chan<- int) {
	for j := range jobsChan {
		fmt.Println("worker", id, "started job", j, "with", len(jobsChan), "jobs left to process")
		time.Sleep(time.Second * 2)
		fmt.Println("worker", id, "finished job", j)
		completedJobsChan <- j
	}

}
