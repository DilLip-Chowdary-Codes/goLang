package main

import "fmt"

func fact(n int) int {
	if n == 1 || n == 0 {
		return 1
	}

	return n * fact(n-1)

}

func fibFunc(nums []int) func() int {
	var fibFunc func(num int) int

	index := -1

	fibFunc = func(num int) int {

		if num < 2 {
			return num
		}
		return fibFunc(num-1) + fibFunc(num-2)
	}

	return func() int {
		index += 1
		num := nums[index]
		return fibFunc(num)

	}
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}

	for _, num := range nums {
		fact := fact(num)
		fmt.Println("Fact of ", num, fact)
	}

	nextFib := fibFunc(nums)

	fmt.Println(nextFib())
	fmt.Println(nextFib())
	fmt.Println(nextFib())
	fmt.Println(nextFib())
	fmt.Println(nextFib())
	fmt.Println(nextFib())
	fmt.Println(nextFib())

}
