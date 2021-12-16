package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	lines := readFile("input")

	passwordsValid := make([]bool, 0)
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		passwordPolicy := parseLine(line)
		passwordsValid = append(passwordsValid, passwordPolicy.isValid())
	}

	validPasswordCount := 0
	for _, valid := range passwordsValid {
		if valid {
			validPasswordCount++
		}
	}
	fmt.Println(validPasswordCount)
}

func readFile(name string) []string {
	file, _ := os.Open(name)
	content, _ := ioutil.ReadAll(file)
	splitContent := strings.Split(string(content), "\n")
	return splitContent
}

func parseLine(line string) passwordPolicy {
	splitLine := strings.Split(line, " ")
	minCharCount, maxCharCount := parsePolicy(splitLine[0])
	targetChar := splitLine[1][0]
	password := splitLine[2]
	return passwordPolicy{minCharCount, maxCharCount, targetChar, password}
}

func parsePolicy(policy string) (int, int) {
	charBound := strings.Split(policy, "-")
	min, _ := strconv.Atoi(charBound[0])
	max, _ := strconv.Atoi(charBound[1])
	return min, max
}

type passwordPolicy struct {
	minCharCount, maxCharCount int
	targetChar                 byte
	password                   string
}

func (policy passwordPolicy) isValid() bool {
	return xor(policy.password[policy.minCharCount-1] == policy.targetChar, policy.password[policy.maxCharCount-1] == policy.targetChar)
}

func xor(a bool, b bool) bool {
	return (a || b) && !(a && b)
}
