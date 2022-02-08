package main

import "fmt"

type studentStruct struct {
	name   string
	age    int
	mobile int64
}

func Student(name string, age int, mobile int64) *studentStruct {
	return &studentStruct{
		name: name, age: age, mobile: mobile,
	}
}

func main() {
	student := Student("DilLip", 23, 9515280160)

	fmt.Println("Student obj", *student)
	fmt.Println("Student obj", student.name)
	fmt.Println("Student obj", &student)
	//fmt.Println("Student obj", student)
	student.age += 1
	fmt.Println("Student obj", student.age)

}
