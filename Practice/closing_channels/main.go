package main

import (
	"fmt"
	"time"
)

func main() {
	notifications := make(chan int, 5)
	done := make(chan string)

	go func() {
		for {
			job, more := <-notifications

			if more {
				fmt.Println("New Notification", job)
			} else {
				fmt.Println("No More Notifications")
				done <- "Done"
				return
			}
		}
	}()

	for i := 0; i < 10; i++ {
		notifications <- i
		fmt.Println("sent job")
		time.Sleep(2 * time.Second)

	}
	close(notifications)

	<-done

}
