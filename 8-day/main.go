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

func main() {

	scanner, file := loadFile("input")
	defer file.Close()

	allCharacters := 0
	allEncodedCharacters := 0
	allWordsCount := 0


	for scanner.Scan() {
		
		line := scanner.Text()
		allCharacters = allCharacters + len(line)
		
		for _, v := range line {

			if v == '"' {
				allEncodedCharacters = allEncodedCharacters + 2		
			}else if v == '\\'{
				allEncodedCharacters = allEncodedCharacters + 2
			}else {
				allEncodedCharacters++
			}
		}
		allEncodedCharacters = allEncodedCharacters + 2

		if(len(line) == 3){
			allWordsCount++
		
		}else{

			stripLine :=  line[1:len(line)-1]

			for i,j:= 0,1; i < len(stripLine); i,j = i+1, j+1 {
				if(stripLine[i] == '\\' && j == len(stripLine)){
					allWordsCount++;
				}else if(stripLine[i] == '\\'){
					switch stripLine[j] {
					case '"':
						allWordsCount++
						i=i+1
						j=i+1
					case '\\':
						allWordsCount++
						i=i+1
						j=i+1
					case 'x':
						allWordsCount++
						i=i+3
						j=i+1
					}
				}else{
					allWordsCount++;
				}
			}
		}
	}
	fmt.Println("1. AllChars - AllStrings = ", allCharacters - allWordsCount)
	fmt.Println("2. AllEncoded - AllChars = ", allEncodedCharacters - allCharacters)

}