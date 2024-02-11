package main

import (
	"bufio"
	"fmt"
	"os"
)

type coord struct {
	x, y int
}

func loadFile(input string) (*bufio.Scanner, *os.File) {
	file, _ := os.Open(input + ".txt")
	return bufio.NewScanner(file), file
}

func RobotAndSanta(chain string) map[coord]bool {
	pathMap := make(map[coord]bool)
	santa, roboSanta := coord{}, coord{}
	pathMap[santa] = true

	for i, rune := range chain {
		c := &santa
		if i%2 == 1 {
			c = &roboSanta
		}

		switch rune {
		case '^':
			c.y++
		case '>':
			c.x++
		case 'v':
			c.y--
		case '<':
			c.x--
		}
		pathMap[*c] = true
	}
	return pathMap

}

func main() {

	scanner, file := loadFile("input")
	defer file.Close()

	for scanner.Scan() {
		fmt.Printf("MultiSanta %d\n", len(RobotAndSanta((scanner.Text()))))
	}
}