package main

import (
	"fmt"
)

func ping(ch chan<- string, msg string, done chan bool) {
	ch <- msg
	done <- true
}

func pong(pings <-chan string, pongs chan<- string, done chan bool) {
	for {
		res, ok := <-pings
		if ok {
			pongs <- res

		} else {
			break
		}
	}
	//
	//for i := 0; i < 5; i++ {
	//	fmt.Println("Num: ", i+1)
	//	time.Sleep(3 * time.Second)
	//}

	done <- true
}

func main() {
	ch1 := make(chan string, 4)
	ch2 := make(chan string, 4)
	done := make(chan bool)

	go ping(ch1, "Ping 1", done)
	<-done
	fmt.Println(len(ch1))

	go ping(ch1, "Ping 2", done)
	<-done
	fmt.Println(len(ch1))

	go ping(ch1, "Ping 3", done)
	<-done
	fmt.Println(len(ch1))

	close(ch1)

	go pong(ch1, ch2, done)
	<-done
	fmt.Println(len(ch1))

	close(ch2)
	for {
		res, isOpen := <-ch2
		if !isOpen {
			break
		}
		fmt.Println(res)

	}

}
