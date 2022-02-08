package main

import (
	"fmt"
)

type animal struct {
	name  string
	color string
}

type dog struct {
	animal
	habits string
}

func (a animal) describe() {
	fmt.Println("Name: ", a.name, "color: ", a.color)
}

func (a dog) describe() {
	fmt.Println("Name: ", a.name, "Color: ", a.color, "Habits: ", a.habits)
}

type describer interface {
	describe()
}

func main() {
	dog := dog{animal{name: "Puppy", color: "Blue"}, "Bark"}

	fmt.Println("Dog: ", dog.name, dog.color, dog.habits)

	dog.describe()

}
