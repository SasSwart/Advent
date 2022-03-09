package main

import (
	"bytes"
	"fmt"
	"os"
)

func input() [][]byte {
	file, _ := os.ReadFile("./input")
	return bytes.Split(file, []byte("\n"))
}

func main() {
	input := input()

	lineLength := len(input[0])
	var bitFrequency []uint = make([]uint, lineLength)
	var gamma, epsilon, mask uint = 0, 0, (1<<lineLength - 1)

	for _, bits := range input {
		for i := 0; i < lineLength; i++ {
			bitFrequency[i] += uint(bits[i] - 48)
		}
	}

	for i := 0; i < len(bitFrequency); i++ {
		gamma = gamma << 1
		gamma += (bitFrequency[i] << 1) / uint(len(input))
	}

	epsilon = mask &^ gamma
	fmt.Printf("%b\n", gamma)
	fmt.Printf("%b\n", epsilon)
	fmt.Printf("%v\n", gamma*epsilon)
}
