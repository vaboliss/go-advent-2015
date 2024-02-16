package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Instruction struct{
	x,y,x1,y1, command int
}

func loadFile(input string) (*bufio.Scanner, *os.File) {
	file, _ := os.Open(input + ".txt")
	return bufio.NewScanner(file), file
}


func ParseFirstCoordinates(instruction *Instruction, coordinates string) { 
	splitCoords := strings.Split(coordinates, ",")

	x, _ := strconv.Atoi(splitCoords[0])
	y, _ := strconv.Atoi(splitCoords[1])

	instruction.x = x
	instruction.y = y

}

func ParseSecondCoordinates(instruction *Instruction, coordinates string){
	splitCoords := strings.Split(coordinates, ",")

	x, _ := strconv.Atoi(splitCoords[0])
	y, _ := strconv.Atoi(splitCoords[1])

	instruction.x1 = x
	instruction.y1 = y
}

func PerformInstructionPart1(grid [][]bool, instruction Instruction) {

	//Turn on
	if(instruction.command == 1){
		for i := instruction.y; i < instruction.y1+1; i++ {
			for j := instruction.x; j < instruction.x1+1; j++ {
				grid[i][j] = true
			}
		}
	// Turn off
	}else if(instruction.command == 2){
		for i := instruction.y; i < instruction.y1+1; i++ {
			for j := instruction.x; j < instruction.x1+1; j++ {
				grid[i][j] = false
			}
		}
	// Toggle
	}else if(instruction.command == 0){
		for i := instruction.y; i < instruction.y1+1; i++ {
			for j := instruction.x; j < instruction.x1+1; j++ {
				grid[i][j] = !grid[i][j]
			}
		}

	}
}

func PerformInstructionPart2(grid [][]int, instruction Instruction) {

	//Turn on
	if(instruction.command == 1){
		for i := instruction.y; i < instruction.y1+1; i++ {
			for j := instruction.x; j < instruction.x1+1; j++ {
				grid[i][j]++
			}
		}
	// Turn off
	}else if(instruction.command == 2){
		for i := instruction.y; i < instruction.y1+1; i++ {
			for j := instruction.x; j < instruction.x1+1; j++ {
				if(grid[i][j] != 0){
					grid[i][j]--
				}
			}
		}
	// Toggle
	}else if(instruction.command == 0){
		for i := instruction.y; i < instruction.y1+1; i++ {
			for j := instruction.x; j < instruction.x1+1; j++ {
				grid[i][j] += 2
			}
		}

	}
}


func main() {

	gridSize := 1000

	part1grid := make([][]bool, gridSize)
	for i := 0; i < gridSize; i++ {
		part1grid[i] = make([]bool, gridSize)
	}

	part2grid := make([][]int, gridSize)
	for i := 0; i < gridSize; i++ {
		part2grid[i] = make([]int, gridSize)
	}

	scanner, file := loadFile("input")
	defer file.Close()


	for scanner.Scan() {
		var instruction Instruction

		line := scanner.Text()
		lineSplit := strings.Split(line, " ")


		if (lineSplit[0] == "toggle"){
			ParseFirstCoordinates(&instruction, lineSplit[1])
			ParseSecondCoordinates(&instruction, lineSplit[3])
			instruction.command = 0


		}else{

			ParseFirstCoordinates(&instruction, lineSplit[2])
			ParseSecondCoordinates(&instruction, lineSplit[4])

			if(lineSplit[1] == "on"){
				instruction.command = 1
			}else{
				instruction.command = 2
			}
		}

		PerformInstructionPart1(part1grid, instruction)
		PerformInstructionPart2(part2grid, instruction)
	}

	counterPart1 := 0
	for i := 0; i < len(part1grid); i++ {
		for j := 0; j < len(part1grid[i]); j++ {
			if part1grid[i][j] {
				counterPart1++
			}		
		}
	}

	counterPart2 := 0
	for i := 0; i < len(part1grid); i++ {
		for j := 0; j < len(part1grid[i]); j++ {
			if(part2grid[i][j] > 0){
				counterPart2 += part2grid[i][j]
			}		
		}
	}

	fmt.Println("First part grid answer: ", counterPart1)
	fmt.Println("Second part grid answer: ", counterPart2)
}