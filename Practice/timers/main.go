package main

import (
	"fmt"
	"time"
)

func main() {

	t1 := time.NewTimer(2 * time.Second)

	<-t1.C

	startTime := time.Now()

	fmt.Println("TImer 1 Done")

	t2 := time.NewTimer(2 * time.Second)

	for {
		fmt.Println("Timer TO Starting")
		<-t2.C
		fmt.Println("Timer 2 Done")
		t2.Reset(2 * time.Second)

		if time.Since(startTime) > 20*time.Second {
			fmt.Println("Request Timed out")
			break
		}
	}

	//t2.Stop()
	//fmt.Println("Timer 2 Stopped")
	//time.Sleep(100 * time.Second)
	fmt.Println("All Done")
}
