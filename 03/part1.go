package main

import (
	"log"
	"os"
)

const inputFile = "input.txt"
const newLine = '\n'
const nonSymbol = '.'
const Nan = 0

type InputArray struct {
	Data [][]byte
}

func (in *InputArray) IsValidNumber(inX, inY int) bool {
	valid := false
	for currX := inX; currX >= 0; currX-- {
		if in.Data[inY][currX] >= '0' && in.Data[inY][currX] <= '9' {
			for y2 := max(0, inY-1); y2 <= min(inY+1, len(in.Data)-1); y2++ {
				for x2 := max(0, currX-1); x2 <= min(currX+1, len(in.Data[y2])-1); x2++ {
					if in.Data[y2][x2] != nonSymbol && (in.Data[y2][x2] < '0' || in.Data[y2][x2] > '9') {
						valid = true
					}
				}
			}
		} else {
			break
		}
	}
	return valid
}

func (in *InputArray) Solve() int {
	total := 0
	currentNumber := Nan
	for y, ySlice := range in.Data {
		currentNumber = Nan
		for x, letter := range ySlice {
			if letter >= '0' && letter <= '9' {
				currentNumber *= 10
				currentNumber += int(letter - '0')
			}
			if letter < '0' || letter > '9' || x == len(ySlice)-1 {
				if currentNumber != Nan {
					if in.IsValidNumber(x-1, y) {
						total += currentNumber
						log.Print(currentNumber)
					}
				}
				currentNumber = Nan
			}
		}
	}

	return total
}

func main() {
	log.Printf("Reading %s", inputFile)
	bytes, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatalf("Cannot read %s", inputFile)
	}

	input := InputArray{Data: make([][]byte, 1)}
	for _, b := range bytes {
		if b == newLine {
			input.Data = append(input.Data, make([]byte, 0))
			continue
		}

		input.Data[len(input.Data)-1] = append(input.Data[len(input.Data)-1], b)
	}

	log.Printf("Total: %d", input.Solve())
}
