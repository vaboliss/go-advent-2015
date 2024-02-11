package main

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type coord struct {
	x, y int
}

func loadFile(input string) (*bufio.Scanner, *os.File) {
	file, _ := os.Open(input + ".txt")
	return bufio.NewScanner(file), file
}

func smallestHashNumbers(secret string, hashStartWith string) int{

	for i := 0; i < 10000000000000; i++ {
		input := secret + strconv.Itoa(i)
		inputInBytes := []byte(input)

		md5Bytes := md5.Sum(inputInBytes)
		md5String := hex.EncodeToString(md5Bytes[:])

		if(strings.HasPrefix(md5String, hashStartWith)){
			return i
		}
	}
	return 0;
}


func main() {

	scanner, file := loadFile("input")
	defer file.Close()

	for scanner.Scan() {
		fmt.Printf("1. Smallest hash numbers that has five 0 in front: %d\n", 
		smallestHashNumbers(scanner.Text(), "00000"))
	
		fmt.Printf("2. Smallest hash numbers that has six 0 in front: %d\n", 
		smallestHashNumbers(scanner.Text(), "000000"))
	
	}
}