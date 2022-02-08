package main

import (
	"fmt"
	"time"
)

func ping(ch chan<- string, msg string) {
	ch <- msg
	//done <- true
}

func ping2(ch chan<- string, msg string) {
	ch <- msg
	//done <- true
}

func pong(pings <-chan string, pongs chan<- string) {
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

	//done <- true
}

func main2() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	//done := make(chan bool)

	go ping(ch1, "Ping 1")
	//<-done
	//fmt.Println(len(ch1))

	go ping(ch1, "Ping 2")
	//<-done
	//fmt.Println(len(ch1))

	go ping(ch1, "Ping 3")
	//<-done
	//fmt.Println(len(ch1))

	go ping(ch1, "Ping 4")

	//ch1 <- "Dummy"

	go ping2(ch2, "Ping 1")
	//<-done
	//fmt.Println(len(ch1))

	go ping2(ch2, "Ping 2")
	//<-done
	//fmt.Println(len(ch1))

	go ping2(ch2, "Ping 3")
	//<-done
	//fmt.Println(len(ch1))

	go ping2(ch2, "Ping 4")

	//ch2 <- "Dummy"
	//close(ch1)
	//for i := 0; i < 1; i++ {
	//	select {
	//	case msg1 := <-ch1:
	//		fmt.Println(msg1)
	//		fmt.Println(len(ch2))
	//
	//		//close(ch1)
	//		//for {
	//		//	res, ok := <-ch1
	//		//	if !ok {
	//		//		break
	//		//	}
	//		//	fmt.Println("Val: ", res)
	//		//
	//		//}
	//
	//	}
	//}

	//go pong(ch1, ch2)
	//<-done
	//fmt.Println(len(ch1))

	//close(ch2)

	for i := 0; i < 2; i++ {
		select {
		case msg2 := <-ch1:
			fmt.Println("Val1: ", msg2)
			fmt.Println(len(ch1))
			//close(ch1)
			for {
				res, ok := <-ch1
				if !ok {
					break
				}
				fmt.Println("Va1: ", res)
				break

			}
		case msg2 := <-ch2:
			fmt.Println("Val2: ", msg2)
			fmt.Println(len(ch2))

			//close(ch2)
			for {
				res, ok := <-ch2
				if !ok {
					break
				}
				fmt.Println("Va1: ", res)
				break

			}
		}

	}

}

func main() {
	c1 := make(chan string)

	c2 := make(chan string)

	startTime := time.Now()

	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "Hello World!"
		//fmt.Println("Channel 1 Done")
	}()

	go func() {
		time.Sleep(3 * time.Second)
		c2 <- "Channel 2 Done"
		//fmt.Println("Channel 2 Done")
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg := <-c1:
			fmt.Println(msg)
		case msg2 := <-c2:
			fmt.Println(msg2)
		}
	}

	timeTaken := time.Since(startTime)

	fmt.Println("Time Taken: ", timeTaken.Seconds(), " Sec")

}
