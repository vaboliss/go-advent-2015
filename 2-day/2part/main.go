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

	totalRibbonPaperRequired := 0

	for scanner.Scan() {
		line := scanner.Text()
		dimentions := strings.Split(line, "x")
		l, _ := strconv.ParseInt(dimentions[0], 0, 64)
		w, _ := strconv.ParseInt(dimentions[1], 0, 64)
		h, _ := strconv.ParseInt(dimentions[2], 0, 64)

		intDimentions := []int{int(l), int(w), int(h)}
		minPerimeter := minPerimeter(intDimentions)

		totalRibbonPaperRequired += minPerimeter
		totalRibbonPaperRequired += int(l * w * h)
	}

	fmt.Println(totalRibbonPaperRequired) // 3737498
}

func minPerimeter(intDimentions []int) int {

	lwPerimeter := 2*intDimentions[0] + 2*intDimentions[1]
	lhPerimeter := 2*intDimentions[0] + 2*intDimentions[2]
	whPerimeter := 2*intDimentions[1] + 2*intDimentions[2]

	return minNumber([]int{lwPerimeter, lhPerimeter, whPerimeter})
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
