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
	for i := range instructions {
		changed, fixedInstructions := fixLine(instructions, i)
		if changed {
			// fmt.Println(fixedInstructions)
			machine := machine{0, 0, fixedInstructions, make(program, 0), getOpcodeDefinitions()}
			fmt.Println(machine.run())
		}
	}
}

type instruction struct {
	line    int
	opcode  string
	address int
}

type program []instruction

type machine struct {
	accumulator       int
	executionPointer  int
	program           program
	stackTrace        program
	opcodeDefinitions map[string]func(int, int, int) (int, int)
}

func (m machine) shouldHalt() bool {
	if m.executionPointer < 0 {
		return true
	}
	if m.executionPointer >= len(m.program) {
		return true
	}
	return false
}

func (m machine) looped() bool {
	for _, instructionInstance := range m.stackTrace {
		if instructionInstance == m.program[m.executionPointer] {
			return true
		}
	}
	return false
}

func (m machine) run() (int, int) {
	for !m.shouldHalt() {
		if m.looped() {
			return -1, m.accumulator
		}
		currentInstruction := m.program[m.executionPointer]
		m.stackTrace = append(m.stackTrace, currentInstruction)
		newAcc, newExecPointer := m.opcodeDefinitions[currentInstruction.opcode](
			m.accumulator,
			m.executionPointer,
			currentInstruction.address)
		m.accumulator = newAcc
		m.executionPointer = newExecPointer
	}
	return 0, m.accumulator
}

func parseInstructions(content string) program {
	instructionLines := strings.Split(string(content), "\n")
	instructions := make(program, 0)
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

func fixLine(p program, i int) (bool, program) {
	np := make(program, 0)
	changed := false

	for lineNumber, instructionInstance := range p {
		if lineNumber != i {
			np = append(np, instructionInstance)
		} else {
			if instructionInstance.opcode == "jmp" {
				changed = true
				np = append(np, instruction{instructionInstance.line, "nop", instructionInstance.address})
			} else if instructionInstance.opcode == "nop" {
				changed = true
				np = append(np, instruction{instructionInstance.line, "jmp", instructionInstance.address})
			} else {
				np = append(np, instructionInstance)
			}
		}
	}
	return changed, np
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
