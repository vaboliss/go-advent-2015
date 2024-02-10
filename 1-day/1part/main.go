package main

import (
	"fmt"
	"os"
)

func main() {
	bytes, err := os.ReadFile("input.txt")

	if err != nil {
		panic(err)
	}

	input := string(bytes)

	floor := 0
	for i := 0; i < len(input); i++ {
		if input[i] == '(' {
			floor++
		} else {
			floor--
		}

	}
	fmt.Println(floor) // Result : 232
}
