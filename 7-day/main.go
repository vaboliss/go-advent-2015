package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type num uint16

type part struct {
	lhs []string
	rhs string
}

type wireVal struct {
	name  string
	value num
}

var binaryPart = map[string]func(inA, inB num) num{
	"AND":    func(inA, inB num) num { return num(inA & inB) },
	"OR":     func(inA, inB num) num { return num(inA | inB) },
	"LSHIFT": func(inA, inB num) num { return num(inA << inB) },
	"RSHIFT": func(inA, inB num) num { return num(inA >> inB) },
}

func getParts(fname string) []part {
	file, _ := os.Open(fname)
	input := bufio.NewScanner(file)
	parts := []part{}
	for input.Scan() {
		line := input.Text()
		tokens := strings.Split(line, " -> ")
		p := part{lhs: strings.Split(tokens[0], " "), rhs: tokens[1]}
		parts = append(parts, p)
	}
	return parts
}

func main() {
	parts := getParts("input.txt")
	// for part 2 I just changed the input in txt file
	wire := map[string]chan num{}
	for _, p := range parts {
		wire[p.rhs] = make(chan num, 100)
	}

	done := make(chan bool)

	wireValues := map[string]num{}
	wireChan := make(chan wireVal)
	go func() {
		for i := 0; i < len(wire); i++ {
			wv := <-wireChan
			wireValues[wv.name] = wv.value
		}
		fmt.Printf("final result: a=%v\n", wireValues["a"])
		done <- true
	}()

	for _, rawP := range parts {
		p := rawP
		lhs, out := p.lhs, p.rhs
		if len(lhs) == 1 {
			n, err := strconv.Atoi(lhs[0])
			if err == nil {
				go func() {
					result := num(n)
					wireChan <- wireVal{out, result}
					wire[out] <- result
				}()
			} else { 
				go func() {
					result := <-wire[lhs[0]]
					wire[lhs[0]] <- result 
					wire[out] <- result
				}()
			}
		} else if len(lhs) == 2 { 
			go func() {
				in := <-wire[lhs[1]]
				wire[lhs[1]] <- in 
				result := ^in      
				wireChan <- wireVal{out, result}
				wire[out] <- result
			}()
		} else { 
			l, op, r := lhs[0], lhs[1], lhs[2]
			fn := binaryPart[op]
			go func() {
				inA, lok := wire[l]
				inB, rok := wire[r]
				var numA, numB num
				if lok && rok {
					numA = <-inA
					inA <- numA 
					numB = <-inB
					inB <- numB 
				} else if lok && !rok {
					b, _ := strconv.Atoi(r)
					numA = <-inA
					inA <- numA 
					numB = num(b)
				} else if !lok && rok {
					a, _ := strconv.Atoi(l)
					numA = num(a)
					numB = <-inB
					inB <- numB 
				} else {
					a, _ := strconv.Atoi(l)
					b, _ := strconv.Atoi(r)
					numA = num(a)
					numB = num(b)
				} 
				result := fn(numA, numB)
				wireChan <- wireVal{out, result}
				wire[out] <- result
			}()
		} 
	} 

	<-done
}