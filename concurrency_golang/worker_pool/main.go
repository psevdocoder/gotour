package main

import (
	"fmt"
	"time"
)

func main() {
	const jobsCount = 15
	const workersCount = 3

	jobs := make(chan int, jobsCount)
	results := make(chan int, jobsCount)
	timeStart := time.Now()

	for i := 0; i < workersCount; i++ {
		go worker(i+1, jobs, results)
	}

	for i := 0; i < jobsCount; i++ {
		jobs <- i + 1
	}

	for i := 0; i < jobsCount; i++ {
		fmt.Println(<-results)
	}
	fmt.Println("time elapsed:", time.Since(timeStart))
}

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		time.Sleep(time.Second) // simulate long-running job
		results <- j * j
		fmt.Println("worker #", id, "finished")
	}
}
