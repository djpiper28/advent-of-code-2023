package main

import (
	"log"
	"os"
)

type ScanLine struct {
	First byte
	Last  byte
}

const inputFile = "input.txt"
const newLine = '\n'

func IsDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func byteToInt(c byte) int {
	return int(c - '0')
}

const notSet = 0

func NewScanLine() ScanLine {
	return ScanLine{First: notSet, Last: notSet}
}

func main() {
	log.Print("Reading", inputFile)
	fileBytes, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatal("Cannot read", inputFile, err)
	}

	total := 0
	currentLine := NewScanLine()
	for _, b := range fileBytes {
		if b == newLine {
			tens := byteToInt(currentLine.First)
			units := byteToInt(currentLine.Last)
			total += tens * 10
			total += units
			// log.Printf("%d%d", tens, units)
			currentLine = NewScanLine()
		}

		if IsDigit(b) {
			currentLine.Last = b
			if currentLine.First == notSet {
				currentLine.First = b
			}
		}
	}

	log.Print("Total", total)
}
