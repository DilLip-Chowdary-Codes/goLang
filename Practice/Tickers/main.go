package main

import (
	"fmt"
	"time"
)

func main() {

	t1 := time.NewTicker(2 * time.Second)

	<-t1.C

	startTime := time.Now()

	fmt.Println("TImer 1 Done")

	//t2 := time.NewTimer(2 * time.Second)

	for {

		timeNow := time.Now()

		<-t1.C

		fmt.Println("Time: ", timeNow.String())

		if time.Since(startTime) > 20*time.Second {
			fmt.Println("Request Timed out")
			break
		}
	}

	startTime = time.Now()

	go func() {
		for {

			timeNow := time.Now()

			<-t1.C

			fmt.Println("Time v2: ", timeNow.String())
		}
	}()

	if time.Since(startTime) > 20*time.Second {
		fmt.Println("Request Timed out")
		t1.Stop()
	}

	//t2.Stop()
	//fmt.Println("Timer 2 Stopped")
	//time.Sleep(100 * time.Second)
	fmt.Println("All Done")
}
