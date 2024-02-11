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


func isStringNiceV2(input string) bool {

	lenghtOfString := len(input)
	for i,j := 0,1; j < lenghtOfString; i,j = i+1, j+1  {

		pair := string(input[i]) + string(input[j])
		substring := input[j:lenghtOfString]

		if(!strings.Contains(substring, pair)){
			continue;
		}

		indexOfSecondPair := strings.Index(substring, pair)

		if(indexOfSecondPair == -1){
			continue
		}

		spaceBetween := input[j:indexOfSecondPair]

		for i := 0; i < len(spaceBetween); i++ {
			
			character := string(spaceBetween[i])
			withoutCharacter := spaceBetween[i+1:len(spaceBetween)-i]
			indexOfSecond := strings.Index(withoutCharacter, character)

			if (indexOfSecond == -1){
				continue
			}

			if(indexOfSecond - i == 1){
				return true
			}


		}

	}

	return false
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

		fmt.Println(stringInput)
		if isStringNice(stringInput){
			firstCounter++
		}

		if isStringNiceV2(stringInput){
			secondCounter++
		}

	}

	fmt.Println("first part number of nice strings:", firstCounter)
	fmt.Println("second part number of nice strings:", secondCounter)
}