package main

import "fmt"

func main(){
	num := 100

	switch num{
	case 1:
		fmt.Println(num,  "is Divisible by 10 or 5")
	case 100:
		fmt.Println(num, "is Divisible by 1")

	default:
		fmt.Println(num, "Default")
	}
}