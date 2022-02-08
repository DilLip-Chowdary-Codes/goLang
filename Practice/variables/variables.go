package main

import (
	"fmt"
	"reflect"
)

func main(){
	var num = 123

	var str string= "string"

	var bool bool = true

	var bool_2 = false

	var float = 3.2

	float_2 := 2.3

	fmt.Println("Num: " , reflect.TypeOf(num))
	fmt.Println("Bool: " , reflect.TypeOf(bool))
	fmt.Println("Str: " , reflect.TypeOf(str))
	fmt.Println("bool_2: " , reflect.TypeOf(bool_2))
	fmt.Println("float: " , reflect.TypeOf(float))
	fmt.Println("float_2: " , reflect.TypeOf(float_2))

}