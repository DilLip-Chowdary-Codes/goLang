package main

import (
	"fmt"
)

type numError struct {
	arg      int
	errorMsg string
}

func (err *numError) Error() string {
	return fmt.Sprintf("arg: %d Error: %s", err.arg, err.errorMsg)

}

func isPositive(num int) (bool, error) {
	if num > 0 {
		return true, nil

	} else if num < 0 {
		return false, nil

	} else {
		return false, &numError{arg: num, errorMsg: "0 Is Neither Positive nor Negative"}
	}
}
func main() {
	nums := []int{1, 2, 3, 4, 0, -1, -3}

	for _, num := range nums {
		if isPositive, err := isPositive(num); isPositive && err == nil {
			fmt.Println("Num: ", num, "Is Positive")
		} else if !isPositive && err == nil {
			fmt.Println("Num: ", num, "Is Negative")
		} else {
			ne := err.(*numError)

			fmt.Println("Error: ", ne.errorMsg)
		}
	}
}
