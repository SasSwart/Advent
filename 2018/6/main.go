package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
)

func main() {
	documents := readFile("input")
	sumAnswers := 0
	for _, d := range documents {
		sumAnswers += d.sumAllYes()
	}
	fmt.Println(sumAnswers)
}

func readFile(name string) []customsForm {
	file, _ := os.Open(name)
	content, _ := ioutil.ReadAll(file)

	delimeter, _ := regexp.Compile("\n\n+")
	documents := delimeter.Split(string(content), -1)

	var parsedDocuments = make([]customsForm, 0)
	for _, doc := range documents {
		parsedDocuments = append(parsedDocuments, newDocument(doc))
	}

	return parsedDocuments
}

func newDocument(text string) customsForm {
	delimeter, _ := regexp.Compile("[ \t\n]+")
	fields := delimeter.Split(text, -1)
	personCount := len(fields)
	answers := make(map[byte]int)
	for _, field := range fields {
		for _, c := range field {
			answers[byte(c)]++
		}
	}
	return customsForm{answers, personCount}
}

func (form customsForm) sumAnyYes() int {
	return len(form.form)
}

func (form customsForm) sumAllYes() int {
	allYes := 0
	for _, answerCount := range form.form {
		if answerCount == form.personCount {
			allYes++
		}
	}
	return allYes
}

type customsForm struct {
	form        map[byte]int
	personCount int
}
