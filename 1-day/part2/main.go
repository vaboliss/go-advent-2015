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
	result := 0
	for i := 0; i < len(input); i++ {
		if input[i] == '(' {
			floor++
		} else {
			floor--
		}

		if floor == -1 {
			result = i + 1
			break
		}

	}

	fmt.Println(result) // Result : 1783
}
