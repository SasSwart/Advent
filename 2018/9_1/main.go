package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	content := readFile("input")
	numbers := parseCipher(content)
	preambleLength := 25
	var invalidNumber int
	for i := range numbers[preambleLength:] {
		if !checkNumberInCipher(numbers, preambleLength, i+preambleLength) {
			invalidNumber = numbers[i+preambleLength]
			break
		}
	}
	fmt.Println(invalidNumber)
}

func readFile(name string) string {
	file, _ := os.Open(name)
	content, _ := ioutil.ReadAll(file)
	return string(content)
}

func parseCipher(content string) []int {
	cipherLines := strings.Split(string(content), "\n")
	numbers := make([]int, 0)
	for _, line := range cipherLines {
		number, _ := strconv.Atoi(line)
		numbers = append(numbers, number)
	}
	return numbers
}

func checkNumberInCipher(numbers []int, preambleLength int, position int) bool {
	numberSet := numbers[position-preambleLength : position]
	target := numbers[position]
	return contains(numberSet, target, 2)
}

func contains(s []int, target int, count int) bool {
	if count == 1 {
		for _, a := range s {
			if a == target {
				return true
			}
		}
		return false
	} else if count > 1 {
		for i, a := range s {
			if contains(s[i:], target-a, count-1) {
				return true
			}
		}
	}
	return false
}
