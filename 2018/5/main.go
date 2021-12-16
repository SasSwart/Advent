package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"sort"
	"strings"
)

func main() {
	seats := readFile("input")
	seatIDs := make([]int, 0)
	for _, seat := range seats {
		row := stringToBinary(seat[0:7], 'F')
		col := stringToBinary(seat[7:10], 'L')
		seatIDs = append(seatIDs, seatID(row, col))
	}
	fmt.Println(findMySeat(seatIDs))
}

func findMySeat(seatIDS []int) int {
	sort.Ints(seatIDS)
	i := int(len(seatIDS) / 2)
	if i == 1 {
		return seatIDS[i] - 1
	}
	if seatIDS[i] != i+seatIDS[0] {
		return findMySeat(seatIDS[:i])
	}
	return findMySeat(seatIDS[i:])
}

func compareSeats(seatList [][]int, i, j int) bool {
	if seatList[i][0] == seatList[j][0] {
		return seatList[i][1] < seatList[j][1]
	}
	return seatList[i][0] < seatList[j][0]
}

func stringToBinary(seat string, char byte) int {
	var seatRow uint = uint(math.Pow(2, float64(len(seat))) - 1)
	var col uint = uint(math.Pow(2, float64(len(seat)-1)))
	for _, c := range seat {
		if c == rune(char) {
			seatRow = (seatRow & ^col)
		}
		col = col >> 1
	}
	return int(seatRow)
}

func seatID(row int, column int) int {
	return row*8 + column
}

func readFile(name string) []string {
	file, _ := os.Open(name)
	content, _ := ioutil.ReadAll(file)
	splitContent := strings.Split(string(content), "\n")
	return splitContent
}
