package main

import (
	"fmt"
	"reflect"
)

func main() {
	var explicit string = "Hello I'm explicit string"
	infered := "Hello I'm infered string"
	fmt.Println("explict string variable type: ", reflect.TypeOf(explicit))
	fmt.Println("Infered string variable type: ", reflect.TypeOf(infered))
}
