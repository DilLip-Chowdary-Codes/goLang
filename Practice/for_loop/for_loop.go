package main

import "fmt"

func main(){

	for i:=1;i<=10;i++{
		fmt.Println(i)
	}

	for i:=10;;i++{

		fmt.Println(i)
		if i == 20 {break}
	}

	for i:=10;;i++{

		if i % 2 != 0 {continue}
		fmt.Println(i)

		if i == 20 {break}
	}
}