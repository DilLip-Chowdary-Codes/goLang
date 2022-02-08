package main

import (
	"fmt"
)

func testFunc(ch chan int) {
	ch <- 100
	ch <- 101
	ch <- 102
	ch <- 103

}

func main() {
	ch := make(chan int)

	go testFunc(ch)
	//time.Sleep(2 * time.Second)
	ch <- 10011
	<-ch
	//ch <- 104
	fmt.Println(len(ch))

	close(ch)
	for {
		res, ok := <-ch
		if ok {
			fmt.Println("OK: ", ok, "Res: ", res)

		} else {
			fmt.Println("OK: ", ok, "Res: ", res)

			break
		}
	}
}
