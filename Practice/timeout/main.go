package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string, 2)

	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "Hello World!"
	}()

	select {
	case msg := <-c1:
		fmt.Println("Channel 1", msg)
	case <-time.After(1 * time.Second):
		fmt.Println("Channel 1 Timed out")

	}

	go func() {
		time.Sleep(10 * time.Second)
		c1 <- "Hello World2!"
	}()

	select {
	case msg := <-c1:
		time.Sleep(2 * time.Second)
		fmt.Println("Channel 1-", msg)
	case <-time.After(5 * time.Second):
		fmt.Println("Channel 1 Timed out2")
	default:
		fmt.Println("No Activity")
	}
}
