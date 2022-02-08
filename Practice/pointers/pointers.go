package main

import "fmt"

func updateVal(num int, newVal int) {
	num = newVal
}

func updateValV2(num *int, newVal int) {
	*num = newVal
}
func updateValWithValInGivenAddress(num int, pointer *int) {
	fmt.Println("Address: ", pointer)
	fmt.Println("Value: ", *pointer)
	num = *pointer
}

func incrementValBy1UsingPointer(pointer *int) {
	*pointer += 1
}

func incrementValBy1UsingVar(val int) {
	val += 1
}

func main() {
	val := 0
	incrementValBy1UsingPointer(&val)
	fmt.Println("Current Val: ", val)
	incrementValBy1UsingVar(val)
	fmt.Println("Current Val: ", val)
	updateVal(val, 10)
	fmt.Println("Current Val: ", val)
	updateValV2(&val, 10)
	fmt.Println("Current Val: ", val)
	newVal := 11
	updateValWithValInGivenAddress(val, &newVal)
	fmt.Println("Current Val: ", val)
	incrementValBy1UsingVar(newVal)
	fmt.Println("Current Val: ", newVal)
	incrementValBy1UsingPointer(&newVal)
	fmt.Println("Current Val: ", newVal)

}
