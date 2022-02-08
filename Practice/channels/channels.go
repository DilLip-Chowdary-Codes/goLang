package main

import (
	"fmt"
	"time"
)

func test(mychnl chan string) {

	mychnl <- "GeeksforGeeks"
	//close(mychnl)
}

func test2(mychnl chan string) {
	mychnl <- "GFG"
	mychnl <- "gfg"
	mychnl <- "Geeks"
}

// Main function
func main() {

	// Creating a channel
	// Using make() function
	mychnl := make(chan string, 100)

	// Anonymous goroutine
	go test(mychnl)

	go test2(mychnl)

	fmt.Println(len(mychnl))

	time.Sleep(2 * time.Second)
	fmt.Println(len(mychnl))

	msg := <-mychnl
	fmt.Println(msg)
	msg = <-mychnl
	fmt.Println(msg)

	msg = <-mychnl
	fmt.Println(msg)

	mychnl <- "New Text"

	// Using for loop
	for {
		res, ok := <-mychnl
		//if  {
		//	fmt.Println("--", res)
		//}
		fmt.Println("OK: ", ok, "Res: ", res)
	}
}
