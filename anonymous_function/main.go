package main

import (
	"fmt"
)

func main() {
	add := func(m int) int {
		return m + 1
	}

	result := add(6)

	fmt.Println(result)
	// fmt.Println(reflect.TypeOf(add))

}
