package main

import "fmt"

type rect struct {
	width  int
	height int
}

func area(r rect) int {
	return r.width * r.height
}

func (r rect) areaV2() int {
	r.height = 10
	return r.width * r.height
}

func perim(r *rect) int {
	return 2 * (r.width + r.height)
}

func main() {
	r := rect{width: 10, height: 5}

	fmt.Println(area(r))
	fmt.Println(perim(&r))
	fmt.Println(r.areaV2())
	fmt.Println("Rect: ", r)
	fmt.Println(perim(&r))
}
