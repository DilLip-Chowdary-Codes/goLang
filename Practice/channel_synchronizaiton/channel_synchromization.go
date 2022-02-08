package main

import (
	"fmt"
	"time"
)

func testFunc(ch chan bool) {
	fmt.Println("Initiated..")

	for i := 0; i < 5; i++ {
		fmt.Println("Num: ", i+1)
		time.Sleep(2 * time.Second)
	}
	ch <- true
}

func main() {
	ch := make(chan bool)

	go testFunc(ch)

	<-ch
	fmt.Println("Done..")
}
