package main

import (
	"fmt"
	"reflect"
)

func main(){
	const num = 123

	const str string= "string"

	const bool bool = true

	const bool_2 = false

	const float = 3.2

	const float_2 := 2.3

	float_2 = 3.5

	fmt.Println("Num: " , reflect.TypeOf(num))
	fmt.Println("Bool: " , reflect.TypeOf(bool))
	fmt.Println("Str: " , reflect.TypeOf(str))
	fmt.Println("bool_2: " , reflect.TypeOf(bool_2))
	fmt.Println("float: " , reflect.TypeOf(float))
	fmt.Println("float_2: " , reflect.TypeOf(float_2))
	fmt.Println("float_2: " , float_2)

}