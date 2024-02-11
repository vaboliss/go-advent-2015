package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	bytes, err := os.ReadFile("input.txt")

	if err != nil {
		panic(err)
	}
	
	input := string(bytes)

	pathMap := make(map[string]int)
	coordinates := []int { 0,0 }
	startingKey := strconv.Itoa(coordinates[0]) + strconv.Itoa(coordinates[1]) 
		
	pathMap[startingKey] = 1
	
	for _, v := range input {
		switch v {
		case '<':
			coordinates[1]-=1
		case '>':
			coordinates[1]+=1
		case '^':
			coordinates[0]+=1
		case 'v':
			coordinates[0]-=1
		}

		pathMap = safeAdd(pathMap, strconv.Itoa(coordinates[0]) + strconv.Itoa(coordinates[1]))

	}
	
	fmt.Println(len(pathMap)) // 2081
}

func safeAdd(pathMap map[string]int, key string) map[string]int {	
	elementAtMap, ok := pathMap[key]

	if ok {
		pathMap[key] = elementAtMap + 1
	}else {
		pathMap[key] = 1
	}

	return pathMap
}