package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	documents := readFile("input")
	validPassports := 0
	for _, d := range documents {
		if d.isValid() {
			validPassports++
		}
	}
	fmt.Println(validPassports)
}

func readFile(name string) []document {
	file, _ := os.Open(name)
	content, _ := ioutil.ReadAll(file)

	delimeter, _ := regexp.Compile("\n\n+")
	documents := delimeter.Split(string(content), -1)

	var parsedDocuments = make([]document, 0)
	for _, doc := range documents {
		parsedDocuments = append(parsedDocuments, newDocument(doc))
	}

	return parsedDocuments
}

func newDocument(text string) document {
	delimeter, _ := regexp.Compile("[ \t\n]+")
	fields := delimeter.Split(text, -1)

	var byr, iyr, eyr, hgt, hcl, ecl, pid, cid string
	for _, field := range fields {
		parsedField := strings.Split(field, ":")
		switch parsedField[0] {
		case "byr":
			byr = parsedField[1]
		case "iyr":
			iyr = parsedField[1]
		case "eyr":
			eyr = parsedField[1]
		case "hgt":
			hgt = parsedField[1]
		case "hcl":
			hcl = parsedField[1]
		case "ecl":
			ecl = parsedField[1]
		case "pid":
			pid = parsedField[1]
		case "cid":
			cid = parsedField[1]
		}
	}
	doc := document{byr, iyr, eyr, hgt, hcl, ecl, pid, cid}
	return doc
}

func (d document) isValid() bool {
	switch "" {
	case d.Byr, d.Iyr, d.Eyr, d.Hgt, d.Hcl, d.Ecl, d.Pid:
		return false
	}

	byr, err := strconv.Atoi(d.Byr)
	if err != nil {
		return false
	}
	if !(1920 <= byr && byr <= 2002) {
		return false
	}

	iyr, err := strconv.Atoi(d.Iyr)
	if err != nil {
		return false
	}
	if !(2010 <= iyr && iyr <= 2020) {
		return false
	}

	eyr, err := strconv.Atoi(d.Eyr)
	if err != nil {
		return false
	}
	if !(2020 <= eyr && eyr <= 2030) {
		return false
	}

	switch d.Hgt[len(d.Hgt)-2:] {
	case "cm":
		hgt, err := strconv.Atoi(d.Hgt[:len(d.Hgt)-2])
		if err != nil {
			return false
		}
		if !(150 <= hgt && hgt <= 193) {
			return false
		}
	case "in":
		hgt, err := strconv.Atoi(d.Hgt[:len(d.Hgt)-2])
		if err != nil {
			return false
		}
		if !(59 <= hgt && hgt <= 76) {
			return false
		}
	default:
		return false
	}

	if d.Hcl[0] != '#' {
		return false
	}
	hclDelimeter, _ := regexp.Compile("^#[0-9a-f]{6}$")
	if hclDelimeter.Match([]byte(d.Hcl)) == false {
		return false
	}
	switch d.Ecl {
	case "amb", "blu", "brn", "gry", "grn", "hzl", "oth":
	default:
		return false
	}

	pidDelimeter, _ := regexp.Compile("^(0+|[1-9])[0-9]{8}$")
	if pidDelimeter.Match([]byte(d.Pid)) == false {
		return false
	}

	return true
}

type document struct {
	Byr string
	Iyr string
	Eyr string
	Hgt string
	Hcl string
	Ecl string
	Pid string
	Cid string
}
