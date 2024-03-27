package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func loadFile(input string) (*bufio.Scanner, *os.File) {
	file, _ := os.Open(input + ".txt")
	return bufio.NewScanner(file), file
}


type AdjCity struct {
	cityName string;
	distance int
}

func main() {

	scanner, file := loadFile("input")
	defer file.Close()

	cityMap := make(map[string][]AdjCity)

	for scanner.Scan() {
		
		line := scanner.Text()
	
		splitLine := strings.Split(line, " ")
	
		cityFrom := splitLine[0]
		cityTo := splitLine[2]
		distance, _ := strconv.Atoi(splitLine[4])

		AddCity(cityMap, cityFrom, cityTo, distance)
		AddCity(cityMap, cityTo, cityFrom, distance)
	}


	shortestNodePath := make([]int, 0)
	longestNodePaths := make([]int, 0)
	for key, _ := range cityMap {
		
		visitedCities := make(map[string]int)
		visitedCities[key] = 1

		shortestDistance := FindShortestPath(key, cityMap, visitedCities)
		shortestNodePath = append(shortestNodePath, shortestDistance)

	}

	for key, _ := range cityMap {
		
		visitedCities := make(map[string]int)
		visitedCities[key] = 1

		longestNode := FindLongestPath(key, cityMap, visitedCities)
		longestNodePaths = append(longestNodePaths, longestNode)

	}


	minDistancePath, _:= findMin(shortestNodePath)
	maxDistancePath, _:= findMax(longestNodePaths)

	fmt.Println("1 part", minDistancePath)
	fmt.Println("2 part", maxDistancePath)
}

func FindShortestPath(key string, cityMap map[string][]AdjCity, visitedCities map[string]int) int {

	city := cityMap[key]
	distances := make([]int, 0)

	for _, value := range city {
		if _, ok := visitedCities[value.cityName]; ok {
			continue;
		
		}else{

			newMap := make(map[string]int)
			for k, v := range visitedCities {
				newMap[k] = v
			}
			visitedCities[value.cityName] = 1

			sdistance := value.distance + FindShortestPath(value.cityName, cityMap, visitedCities)
			
			visitedCities = newMap
			distances = append(distances, sdistance)
		}
	}

	min, _ := findMin(distances)

	return min;
}

func FindLongestPath(key string, cityMap map[string][]AdjCity, visitedCities map[string]int) int {

	city := cityMap[key]
	distances := make([]int, 0)

	for _, value := range city {
		if _, ok := visitedCities[value.cityName]; ok {
			continue;
		
		}else{

			newMap := make(map[string]int)
			for k, v := range visitedCities {
				newMap[k] = v
			}
			visitedCities[value.cityName] = 1

			sdistance := value.distance + FindLongestPath(value.cityName, cityMap, visitedCities)
			
			visitedCities = newMap
			distances = append(distances, sdistance)
		}
	}

	min, _ := findMax(distances)

	return min;
}

func AddCity(cityMap map[string][]AdjCity, cityFrom string, cityTo string, distance int) {
	if existingArray, ok := cityMap[cityFrom]; ok {
		myArray := append(existingArray, AdjCity{cityName: cityTo, distance: distance})
		cityMap[cityFrom] = myArray
	} else {
		adjCities := []AdjCity{AdjCity{cityName: cityTo, distance: distance}}
		cityMap[cityFrom] = adjCities
	}
}

func findMin(arr []int) (int, error) {
	if len(arr) == 0 {
	  return 0, errors.New("Cannot find minimum in empty array")
	}
  
	min := arr[0]
	for _, value := range arr {
	  if value < min {
		min = value
	  }
	}

	return min, nil
  }


func findMax(arr []int) (int, error) {
	if len(arr) == 0 {
	  return 0, errors.New("Cannot find max in empty array")
	}
  
	max := arr[0]
	for _, value := range arr {
	  if value > max {
		max = value
	  }
	}

	return max, nil
  }