package main

import (
	"fmt"
	"time"
)

func printHello(src string) {
	fmt.Println("HelloWorld From ", src)

	for i := 0; i < 5; i++ {

		fmt.Println(src, i)
	}
}
func main() {
	printHello("Normal")
	go printHello("Go Routines")

	time.Sleep(1 * time.Second)
	go func(src string) {
		fmt.Println("HelloWorld From ", src)

		for i := 0; i < 5; i++ {

			fmt.Println(src, i)
		}
	}("Anonymous")

	time.Sleep(1 * time.Second)
	fmt.Println("Hello From Local")
}
