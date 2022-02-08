package main

import (
	"fmt"
)

func main() {

	s := make([]string, 0)

	s = append(s, "a")
	s = append(s, "b")
	s = append(s, "c")
	s = append(s, "d")
	s = append(s, "e")

	fmt.Println(s)

	s2 := make([]string, len(s))

	copy(s2, s)

	fmt.Println(s2[2:4])

	s3 := []string{"1", "2", "3"}

	arr3 := []int{11, 13, 4, 5}

	s3 = append(s3, "34")

	s5 := make([]int, 0)

	s5 = append(arr3, 34)

	fmt.Println(s3)
	fmt.Println(arr3)
	fmt.Println(s5)

	fmt.Printf("%T", s3)
	fmt.Printf("%T", s5)
	fmt.Printf("%T", arr3)

	fmt.Println("Multi Dimensional Slice")

	s6 := make([][]int, 3)

	for i := 0; i < 3; i++ {
		lenOfSlice := i + 1

		s6[i] = make([]int, lenOfSlice)

		for j := 0; j < lenOfSlice; j++ {
			s6[i][j] = i + j
		}
	}
	fmt.Println("S4: ", s6)
}
