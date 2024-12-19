package main

import (
	"fmt"
	"io"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// Utils
func readFile() ([]int, []int) {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	content, _ := io.ReadAll(file)
	lines := strings.Split(string(content), "\n\n")

	matches := regexp.MustCompile(`\d+`).FindAllStringSubmatch(lines[0], -1)
	var registers []int
	for _, match := range matches {
		number, _ := strconv.Atoi(match[0])
		registers = append(registers, number)
	}

	instructions := regexp.MustCompile(`\d+`).FindAllStringSubmatch(lines[1], -1)
	var instructionsList []int
	for _, instruction := range instructions {
		number, _ := strconv.Atoi(instruction[0])
		instructionsList = append(instructionsList, number)
	}

	return instructionsList, registers
}

func getLiteralValue(registers []int, literal int) int {
	if literal <= 3 {
		return literal
	}
	if literal == 4 {
		return registers[0]
	}
	if literal == 5 {
		return registers[1]
	}
	if literal == 6 {
		return registers[2]
	}
	panic("Invalid literal")
}

func executeOperation(opcode int, registers []int, literal int, instructionPointer int, outputs []int) (int, []int, []int){
	if opcode == 0 {
		comboOperand := getLiteralValue(registers, literal)
		numerator := registers[0]
		denominator := int(math.Pow(2, float64(comboOperand)))
		result := numerator / denominator
		registers[0] = result
	} else if opcode == 1 {
		registers[1] = registers[1] ^ literal
	} else if opcode == 2 {
		comboOperand := getLiteralValue(registers, literal)
		registers[1] = comboOperand % 8
	} else if opcode == 3 {
		if registers[0] != 0 {
			return literal, registers, outputs
		}
	} else if opcode == 4 {
		registers[1] = registers[1] ^ registers[2]
	} else if opcode == 5 {
		comboOperand := getLiteralValue(registers, literal)
		outputs = append(outputs, comboOperand % 8)
	} else if opcode == 6 {
		comboOperand := getLiteralValue(registers, literal)
		numerator := registers[0]
		denominator := int(math.Pow(2, float64(comboOperand)))
		result := numerator / denominator
		registers[1] = result
	} else if opcode == 7 {
		comboOperand := getLiteralValue(registers, literal)
		numerator := registers[0]
		denominator := int(math.Pow(2, float64(comboOperand)))
		result := numerator / denominator
		registers[2] = result
	}

	return instructionPointer + 2, registers, outputs
}

func intArrayToString(array []int) string {
	return strings.Join(strings.Fields(fmt.Sprint(array)), ",")
}

// Part 1
func executeOperations(registers []int, instructions []int) []int {
	instructionPointer := 0
	outputs := []int{}

	for instructionPointer < len(instructions) {
		opcode := instructions[instructionPointer]
		literal := instructions[instructionPointer+1]
		instructionPointer, registers, outputs = executeOperation(opcode, registers, literal, instructionPointer, outputs)
	}

	return outputs
}

func part1(){
	instructions, registers := readFile()
	results := executeOperations(registers, instructions)
	fmt.Println(intArrayToString(results))
}

// Part 2
func areAllEqual(smallArray []int, bigArray []int) bool {
	for i, value := range smallArray {
		if value != bigArray[i] {
			return false
		}
	}

	return true
}

func executeOperationsAndCheckOutput(registers []int, instructions []int, expectedOutput []int, expectedOutputAsString string) bool {
	instructionPointer := 0
	outputs := []int{}

	for instructionPointer < len(instructions) {
		opcode := instructions[instructionPointer]
		literal := instructions[instructionPointer+1]
		instructionPointer, registers, outputs = executeOperation(opcode, registers, literal, instructionPointer, outputs)
		if len(outputs) > len(expectedOutput) {
			return false
		}
		if len(outputs) == len(expectedOutput) {
			break
		}
		if !areAllEqual(outputs, expectedOutput) {
			return false
		}
	}

	return intArrayToString(outputs) == expectedOutputAsString
}

func copyRegisters(registers []int) []int {
	var newRegisters []int
	for _, register := range registers {
		newRegisters = append(newRegisters, register)
	}
	return newRegisters
}

func reverseArray(array []int) []int {
	var reversedArray []int

	for i := len(array) - 1; i >= 0; i-- {
		reversedArray = append(reversedArray, array[i])
	}

	return reversedArray
}

func findOutput(registers []int, instructions []int, expectedOutput int, expectedOutputAsString string) int {
	fmt.Println(expectedOutput)
	i := 0
	for {
		newRegisters := copyRegisters(registers)
		newRegisters[0] = i

		solved := executeOperationsAndCheckOutput(newRegisters, instructions, []int{expectedOutput}, expectedOutputAsString)
		if solved {
			return i
		}
		i++
	}
}

// Part 2
func findSolution(instructions []int, answer int) int{
	if len(instructions) == 0 {
		return answer
	}

	// Hardcoded operations from my input
	for i := 0; i<8; i++ {
		a := (answer << 3) + i
		b := a % 8
		b = b ^ 7
		c := a / int(math.Pow(2, float64(b)))
		b = b ^ 7
		b = b ^ c
		if b % 8 == instructions[len(instructions)-1] {
			solution := findSolution(instructions[:len(instructions)-1], a)
			if solution != -1 {
				return solution
			}
		}
	}

	return -1
}

func part2(){
	instructions, _ := readFile()
	i := findSolution(instructions, 0)

	fmt.Println(i)
}

func main() {
	part1()
	part2()
}