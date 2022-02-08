package main

import (
	"fmt"
	"sync"
	"time"
)

func f1(wg *sync.WaitGroup) {
	time.Sleep(3 * time.Second)
	fmt.Println("Testing 1")
	wg.Done()
}

func f2(wg *sync.WaitGroup) {
	time.Sleep(3 * time.Second)
	fmt.Println("Testing 2")
	wg.Done()

}

func main() {
	wg := sync.WaitGroup{}

	wg.Add(1)
	go f1(&wg)
	wg.Add(1)
	go f2(&wg)

	wg.Wait()
}
