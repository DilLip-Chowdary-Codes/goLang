package main

import (
	"fmt"
	"strconv"
)

func main() {
	map1 := make(map[string]int)
	for i := 0; i < 5; i++ {
		map1[strconv.FormatInt(int64(i), 10)] = i
	}
	fmt.Println(map1)

	fmt.Println(len(map1))

	delete(map1, "1")

	fmt.Println(map1)

	fmt.Println(len(map1))

	map2 := map[string]int{"1": 1}

	fmt.Println(map2)

}
