package main

import (
	"fmt"
	"log"
	"os"
)

const inputFile = "input.txt"
const newLine = '\n'
const nonSymbol = '.'
const gearSymbol = '*'
const Nan = 0

type InputArray struct {
	Data      [][]byte
	SeenGears map[string]bool
}

func (in *InputArray) ScanNumber(x, y int) (int, int) {
	if in.Data[y][x] < '0' || in.Data[y][x] > '9' {
		return 0, -1
	}

	startIndex := x
	for startIndex >= 0 {
		if in.Data[y][startIndex] < '0' || in.Data[y][startIndex] > '9' {
			// Overscanned by one
			startIndex++
			break
		}
		startIndex--
	}
	startIndex = max(0, startIndex)

	number := 0
	for x3 := startIndex; x3 < len(in.Data[y]); x3++ {
		if in.Data[y][x3] < '0' || in.Data[y][x3] > '9' {
			break
		}
		number *= 10
		number += int(in.Data[y][x3] - '0')
	}
	return number, startIndex + y*len(in.Data)
}

func (in *InputArray) HandleGear(x, y int) uint64 {
	first := 0
	firstIndex := -1
	last := 0
	lastIndex := -1

	for y2 := max(0, y-1); y2 <= min(y+1, len(in.Data)-1); y2++ {
		for x2 := max(0, x-1); x2 <= min(x+1, len(in.Data[y])-1); x2++ {
			num, index := in.ScanNumber(x2, y2)
			if index == -1 {
				continue
			}

			if firstIndex == -1 {
				first = num
				firstIndex = index
			} else if firstIndex != index {
				last = num
				lastIndex = index
			}
		}
	}

	if firstIndex == -1 || lastIndex == -1 {
		log.Print("Lone number")
		return 0
	}

	var total uint64 = uint64(first * last)
	log.Printf("Found a gear at %d, %d; total %d", x, y, total)
	log.Printf("Gear numbers: %d, %d", first, last)
	return total
}

func (in *InputArray) ScanForGear(inX, inY int) uint64 {
	var total uint64 = 0
	for y2 := max(0, inY-1); y2 <= min(inY+1, len(in.Data)-1); y2++ {
		for x2 := max(0, inX-1); x2 <= min(inX+1, len(in.Data[y2])-1); x2++ {
			if in.Data[y2][x2] == gearSymbol {
				gearName := fmt.Sprintf("%d,%d", x2, y2)
				found, _ := in.SeenGears[gearName]
				if found {
					log.Printf("Duplicate gear detected at %d, %d", x2, y2)
				} else {
					total += in.HandleGear(x2, y2)
					in.SeenGears[gearName] = true
				}
			}
		}
	}
	return total
}

func (in *InputArray) HandleNumber(inX, inY int) uint64 {
	var total uint64 = 0
	for currX := inX; currX >= 0; currX-- {
		if in.Data[inY][currX] >= '0' && in.Data[inY][currX] <= '9' {
			total += in.ScanForGear(currX, inY)
		} else {
			break
		}
	}
	return total
}

func (in *InputArray) Solve() uint64 {
	var total uint64 = 0
	currentNumber := Nan
	for y, ySlice := range in.Data {
		log.Printf("Solving for y %d/%d", y+1, len(in.Data))
		currentNumber = Nan
		for x, letter := range ySlice {
			if letter >= '0' && letter <= '9' {
				currentNumber *= 10
				currentNumber += int(letter - '0')
			}
			if letter < '0' || letter > '9' || x == len(ySlice)-1 {
				if currentNumber != Nan {
					total += in.HandleNumber(x-1, y)
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

	input := InputArray{Data: make([][]byte, 1), SeenGears: make(map[string]bool)}
	for _, b := range bytes {
		if b == newLine {
			input.Data = append(input.Data, make([]byte, 0))
			continue
		}

		input.Data[len(input.Data)-1] = append(input.Data[len(input.Data)-1], b)
	}

	log.Print("Solving")
	log.Printf("Total: %d", input.Solve())
}
