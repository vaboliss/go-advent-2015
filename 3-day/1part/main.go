package main

import (
	"bufio"
	"os"
)

func main() {
	inFile, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer inFile.Close()

	scanner := bufio.NewScanner(inFile)
	for scanner.Scan() {
	}

}