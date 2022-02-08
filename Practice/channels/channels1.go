package main

import (
	"fmt"
	"time"
)

func newFunc(ch chan int) {
	num := 0
	fmt.Println("Num: ", num)
	for i := 0; i < 10; i++ {
		//fmt.Println(num, i)
	}

	time.Sleep(2 * time.Second)
	ch <- 100
	ch <- 101
	ch <- 102

	//close(ch)
}

func main() {
	channel2 := make(chan int, 4)

	//go newFunc(channel2)
	fmt.Println("Len: ", len(channel2))
	fmt.Println("Waiting")
	fmt.Println("Len: ", len(channel2))
	channel2 <- 20
	channel2 <- 21
	channel2 <- 22

	//for nums := range channel2 {
	//	fmt.Println("Channel Nums: ", nums)
	//}
	fmt.Println("Len: ", len(channel2))

	close(channel2)
	fmt.Println("Len: ", len(channel2))

	for {
		res, ok := <-channel2
		if ok == false {
			fmt.Println("Channel Close ", ok)
			break
		}
		fmt.Println("Channel Open ", res, ok)
	}

	fmt.Println("End Main method")
}
