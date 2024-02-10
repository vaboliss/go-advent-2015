package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	inFile, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer inFile.Close()

	scanner := bufio.NewScanner(inFile)

	totalWrappingPaperRequired := 0

	for scanner.Scan() {
		line := scanner.Text()
		dimentions := strings.Split(line, "x")
		l, _ := strconv.ParseInt(dimentions[0], 0, 64)
		w, _ := strconv.ParseInt(dimentions[1], 0, 64)
		h, _ := strconv.ParseInt(dimentions[2], 0, 64)

		wrappingPaperForPresent := 2*l*w + 2*w*h + 2*h*l

		totalWrappingPaperRequired += int(wrappingPaperForPresent)
		totalWrappingPaperRequired += minNumber([]int{int(l * w), int(w * h), int(h * l)})

	}

	fmt.Println(totalWrappingPaperRequired) // 1586300
}

func minNumber(arr []int) int {
	min := arr[0]

	for i := 1; i < len(arr); i++ {
		if min > arr[i] {
			min = arr[i]
		}
	}

	return min
}
