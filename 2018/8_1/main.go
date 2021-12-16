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
	instructions := parseInstructions(content)
	opcodeDefinitions := getOpcodeDefinitions()
	machine := machine{0, 0, instructions, make([]instruction, 0), opcodeDefinitions}
	fmt.Println(machine.run())
}

func parseInstructions(content string) []instruction {
	instructionLines := strings.Split(string(content), "\n")
	instructions := make([]instruction, 0)
	for i, line := range instructionLines {
		instructions = append(instructions, newInstruction(i, line))
	}
	return instructions
}

func readFile(name string) string {
	file, _ := os.Open(name)
	content, _ := ioutil.ReadAll(file)
	return string(content)
}

func newInstruction(line int, instructionLine string) instruction {
	instructionComponents := strings.Split(string(instructionLine), " ")
	opcode := instructionComponents[0]
	address, _ := strconv.Atoi(instructionComponents[1])
	return instruction{line, opcode, address}
}

type instruction struct {
	line    int
	opcode  string
	address int
}

type machine struct {
	accumulator       int
	executionPointer  int
	program           []instruction
	stackTrace        []instruction
	opcodeDefinitions map[string]func(int, int, int) (int, int)
}

func newMachine(program []instruction) {

}

func (m machine) shouldHalt() bool {
	for _, instructionInstance := range m.stackTrace {
		if instructionInstance == m.program[m.executionPointer] {
			return true
		}
	}
	return false
}

func (m machine) run() int {
	for !m.shouldHalt() {
		currentInstruction := m.program[m.executionPointer]
		m.stackTrace = append(m.stackTrace, currentInstruction)
		fmt.Println(currentInstruction)
		newAcc, newExecPointer := m.opcodeDefinitions[currentInstruction.opcode](m.accumulator, m.executionPointer, currentInstruction.address)
		m.accumulator = newAcc
		m.executionPointer = newExecPointer
	}
	return m.accumulator
}

func getOpcodeDefinitions() map[string]func(int, int, int) (int, int) {
	opcodeDefinitions := make(map[string]func(int, int, int) (int, int))
	opcodeDefinitions["acc"] = func(acc, pointer, arg int) (int, int) {
		return acc + arg, pointer + 1
	}
	opcodeDefinitions["jmp"] = func(acc, pointer, arg int) (int, int) {
		return acc, pointer + arg
	}
	opcodeDefinitions["nop"] = func(acc, pointer, arg int) (int, int) {
		return acc, pointer + 1
	}
	return opcodeDefinitions
}
