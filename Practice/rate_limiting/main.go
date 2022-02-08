package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan time.Time, 5)

	tick := time.NewTicker(200 * time.Millisecond)

	for i := 0; i < 5; i++ {
		ch1 <- time.Now()
	}

	close(ch1)

	for val := range ch1 {
		<-tick.C
		fmt.Println("request", val.String()[0], time.Now())
	}

	burstyLimiter := make(chan time.Time, 3)

	println("Bursty Limiter...")

	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	go func() {
		for t := range time.Tick(2 * time.Second) {
			fmt.Println("Range t: ", t)
			burstyLimiter <- t
			fmt.Println("Test")
		}
	}()

	burstyRequests := make(chan int, 6)

	for i := 1; i <= 6; i++ {
		burstyRequests <- i
	}

	close(burstyRequests)

	for req := range burstyRequests {
		fmt.Println("Bursty Limiter Length: ", len(burstyLimiter))

		<-burstyLimiter
		fmt.Println("request", req, time.Now())

		time.Sleep(2 * time.Second)

	}

	fmt.Println("TIme : ", time.Now())

	time.Sleep(20 * time.Second)
	fmt.Println("End TIme : ", time.Now())

}
