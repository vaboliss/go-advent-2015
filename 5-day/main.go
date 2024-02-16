package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type coord struct {
	x, y int
}

func loadFile(input string) (*bufio.Scanner, *os.File) {
	file, _ := os.Open(input + ".txt")
	return bufio.NewScanner(file), file
}


func isStringNiceV2(input string) int {

	var nice int
	passesRule1 := func(line string) bool {
		for i := 0; i < len(line)-2; i++ {
			toMatch := line[i : i+2]
			for j := i + 2; j < len(line)-1; j++ {
				if line[j:j+2] == toMatch {
					return true
				}
			}
		}
		return false
	}

	for _, line := range strings.Split(input, "\n") {
		rule1 := passesRule1(line)

		var rule2 bool
		for i := 0; i < len(line)-2; i++ {
			if line[i] == line[i+2] {
				rule2 = true
				break
			}
		}
		if rule1 && rule2 {
			nice++
		}
	}

	return nice
}

func isStringNice(input string) bool {

	if(strings.Contains(input, "ab") || 
	 	strings.Contains(input, "cd") || 
		strings.Contains(input, "pq") ||
		strings.Contains(input, "xy")){
		
		return false
	}

	if(!strings.ContainsAny(input, "aeiou")){
		
		return false
	}

	requiredNumbersOfVowels := 3
	threeVowelsExist :=false
	twoLettersInTheRow := false

	var lastLetter byte;

	for i := 0; i < len(input); i++ {
		
		if(!twoLettersInTheRow &&
			lastLetter == input[i]){
			twoLettersInTheRow = true
		}
		lastLetter = input[i]

		if(!threeVowelsExist &&
			strings.Contains("aeiou", string(input[i]))){
			requiredNumbersOfVowels--
			
			if(requiredNumbersOfVowels == 0){
				threeVowelsExist = true
			}
		}

		if(twoLettersInTheRow && threeVowelsExist){
			return true;
		}

	}

	return false
}


func main() {

	scanner, file := loadFile("input")
	defer file.Close()

	firstCounter := 0;
	secondCounter := 0;

	for scanner.Scan() {
		stringInput := scanner.Text()

		if isStringNice(stringInput){
			firstCounter++
		}

	}
	 
	allTextBytes, _ := os.ReadFile("input.txt")
	secondCounter = isStringNiceV2(string(allTextBytes));

	fmt.Println("first part number of nice strings:", firstCounter)
	fmt.Println("second part number of nice strings:", secondCounter)
}