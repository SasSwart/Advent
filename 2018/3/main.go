package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	lines := readFile("input")
	deltas := [...][2]int{{1, 1}, {3, 1}, {5, 1}, {7, 1}, {1, 2}}

	product := 1
	for _, delta := range deltas {
		product *= countTrees(lines, delta)
	}
	fmt.Println(product)
}

func countTrees(lines []string, slope [2]int) int {
	isTree := make([]bool, 0)
	for i := 0; i*slope[1] < len(lines); i++ {
		y := i * slope[1]
		x := (i * slope[0]) % len(lines[y])
		isTree = append(isTree, lines[y][x] == '#')
	}

	treeCounter := 0
	for _, v := range isTree {
		if v {
			treeCounter++
		}
	}

	return treeCounter
}

func readFile(name string) []string {
	file, _ := os.Open(name)
	content, _ := ioutil.ReadAll(file)
	splitContent := strings.Split(string(content), "\n")
	return splitContent
}
