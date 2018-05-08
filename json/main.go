package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	packet := "packet"

	jsonPacket, ok := json.Marshal(packet)
	if ok != nil {
		panic("Could not Marshal object")
	}
	fmt.Println(string(jsonPacket))
}
