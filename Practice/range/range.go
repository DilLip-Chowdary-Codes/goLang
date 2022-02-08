package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4}

	for i := range nums {
		fmt.Println("index: ", i, "val: ", i)
	}

	maps := map[string]int{"1": 1, "two": 2}

	for i, v := range maps {
		fmt.Println("index: ", i, "val: ", v)
	}

	for i := range maps {
		fmt.Println("index: ", i, "val: ", i)
	}

	for i, v := range "go Lang" {
		fmt.Println("index: ", i, "val: ", v)
	}
}
