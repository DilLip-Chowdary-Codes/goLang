package main

import "fmt"

func is_even(nums ...int) map[int]map[string]bool {
	result := make(map[int]map[string]bool, 0)

	for _, num := range nums {
		if num%2 == 0 {
			result[num] = map[string]bool{"isEven": true, "isOdd": false}

		} else {
			result[num] = map[string]bool{"isEven": false, "isOdd": true}
		}
	}

	return result
}

func main() {
	nums := []int{1, 2, 3, 4}

	result := is_even(nums...)

	fmt.Println(result)
}
