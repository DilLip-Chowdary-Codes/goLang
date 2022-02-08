package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var counter uint64

	wg := sync.WaitGroup{}

	for n := 0; n < 50; n++ {
		wg.Add(1)

		go func() {

			for i := 0; i < 1000; i++ {
				atomic.AddUint64(&counter, 1)
			}
			fmt.Println("Current Val: ", atomic.LoadUint64(&counter))
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println("Counter Val: ", counter)

}
