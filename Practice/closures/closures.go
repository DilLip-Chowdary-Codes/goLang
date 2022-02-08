package main

import "fmt"

func getPerson(people []string) func() string {
	currentIndex := -1

	return func() string {
		currentIndex++
		return people[currentIndex]
	}
}
func main() {
	people := []string{"DilLip", "Raja", "Tharun", "Mouni", "Kavya"}

	getPerson := getPerson(people)

	fmt.Println("Order by Distance not by Love")

	fmt.Println(getPerson())
	fmt.Println(getPerson())
	fmt.Println(getPerson())
	fmt.Println(getPerson())

}
