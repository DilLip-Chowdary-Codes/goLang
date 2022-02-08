package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	//println("Worker Called", id)

	for j := range jobs {
		fmt.Println("worker", id, "started job", j)
		time.Sleep(3 * time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

func main() {
	jobs := make(chan int, 10)
	results := make(chan int, 10)

	startTime := time.Now().UTC()

	fmt.Println(startTime.String())

	for i := 1; i < 5; i++ {
		go worker(i, jobs, results)
	}

	for j := 1; j <= 10; j++ {
		jobs <- j
	}

	//close(jobs)

	for a := 1; a <= 10; a++ {
		msg := <-results

		fmt.Println("Result ", msg)
	}

	timeTaken := time.Since(startTime).String()

	println("Time Taken: ", timeTaken[:len(timeTaken)-1], " Sec")
}
