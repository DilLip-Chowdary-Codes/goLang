package main

import "fmt"

func main(){

	arr_1:= [5]int {1, 2, 3, 4, 5}

	// s := make([]string, 3)

	fmt.Println(arr_1)

	var arr_2 [5] int

	arr_2[2] = arr_1[3]

	fmt.Println(arr_2)

	arr_3 := [4]int{11, 12, 13, 4, 5}

	fmt.Println(arr_3)
}