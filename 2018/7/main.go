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
	content := readFile("input")
	mayBeContained := parseRulesContained(content)

	numBags := 0
	for _, col := range buildTree(mayBeContained, bagCollection{"shiny gold", 1}).buildNodeSet() {
		numBags += col
	}
	fmt.Println(numBags - 1)
}

func buildTree(graph map[string][]bagCollection, seed bagCollection) node {
	children := make([]node, 0)
	for _, bagCollectionNode := range graph[seed.color] {
		children = append(children, buildTree(graph, bagCollectionNode))
	}
	return node{seed, children}
}

func getContainerBags(rules map[string]map[string]bool, frontier map[string]bool) map[string]bool {
	for bag := range frontier {
		for newFrontierBag := range rules[bag] {
			frontier[newFrontierBag] = true
		}
	}
	return frontier
}

func readFile(name string) string {
	file, _ := os.Open(name)
	content, _ := ioutil.ReadAll(file)
	return string(content)
}

func parseRulesContain(content string) map[string][]bagCollection {
	splitContent := strings.Split(string(content), "\n")
	mayBeContainedRules := make(map[string][]bagCollection)

	for _, ruleLine := range splitContent {
		container, containees := parseRule(ruleLine)
		for _, containee := range containees {
			mayBeContainedRules[containee.color] = append(mayBeContainedRules[containee.color], bagCollection{container, containee.count})
		}
	}

	return mayBeContainedRules
}

func parseRulesContained(content string) map[string][]bagCollection {
	splitContent := strings.Split(string(content), "\n")
	mayBeContainedRules := make(map[string][]bagCollection)

	for _, ruleLine := range splitContent {
		container, containees := parseRule(ruleLine)
		for _, bagCollection := range containees {
			mayBeContainedRules[container] = append(mayBeContainedRules[container], bagCollection)
		}
	}

	return mayBeContainedRules
}

func parseRule(rule string) (string, []bagCollection) {
	rulePattern, _ := regexp.Compile("([\\w ]+) bags contain (((\\d+) (([\\w ]+)bags?,?)+)|(no other bags))\\.")
	containeeDelimeter, _ := regexp.Compile(" bags?,? ?")
	containeePattern, _ := regexp.Compile("(\\d+) ([a-z ]+)")

	match := rulePattern.FindAllStringSubmatch(rule, -1)[0]

	container := match[1]
	containees := make([]bagCollection, 0)

	bagMayContainOtherBags := match[7] == ""
	if bagMayContainOtherBags {
		containeeStrings := containeeDelimeter.Split(match[3], -1)
		for _, bagString := range containeeStrings {
			parsedContaineeString := containeePattern.FindAllStringSubmatch(bagString, -1)

			if len(parsedContaineeString) > 0 {
				count, _ := strconv.Atoi(parsedContaineeString[0][1])
				color := parsedContaineeString[0][2]
				containees = append(containees, bagCollection{color, count})
			}
		}
	}
	return container, containees
}

type node struct {
	bags     bagCollection
	children []node
}

func (root node) buildNodeSet() map[string]int {
	nodeSet := make(map[string]int)
	nodeSet[root.bags.color] = root.bags.count
	for _, child := range root.children {
		childNoteSet := child.buildNodeSet()
		for key, count := range childNoteSet {
			nodeSet[key] += count * root.bags.count
		}
	}
	return nodeSet
}

type bagCollection struct {
	color string
	count int
}
