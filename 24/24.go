package main

import (
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Utils
type LogicGate struct {
	input1 string
	input2 string
	operation string
}

func readFile() (map[string]int, map[string]LogicGate) {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	content, _ := io.ReadAll(file)
	contentAsString := string(content)

	wireData := make(map[string]int)
	blocks := strings.Split(contentAsString, "\n\n")
	for _, block := range strings.Split(blocks[0], "\n") {
		parts := strings.Split(block, ": ")
		value, _ := strconv.Atoi(parts[1])
		wireData[parts[0]] = value
	}

	logicGates := make(map[string]LogicGate)
	for _, block := range strings.Split(blocks[1], "\n") {
		elems := strings.Split(block, " ")
		input1 := elems[0]
		operation := elems[1]
		input2 := elems[2]
		output := elems[4]
		logicGates[output] = LogicGate{input1, input2, operation}
		wireData[output] = -1
	}

	return wireData, logicGates
}

func areAllValuesAvailable(wires map[string]int) bool {
	for _, value := range wires {
		if value == -1 {
			return false
		}
	}
	return true
}

func calculateAvailableValues(wires map[string]int, logicGates map[string]LogicGate) map[string]int {
	for !areAllValuesAvailable(wires) {
		for wire, logicGate := range logicGates {
			input1 := wires[logicGate.input1]
			if input1 == -1 {
				continue
			}

			input2 := wires[logicGate.input2]
			if input2 == -1 {
				continue
			}

			if logicGate.operation == "AND" {
				wires[wire] = input1 & input2
			}

			if logicGate.operation == "OR" {
				wires[wire] = input1 | input2
			}

			if logicGate.operation == "XOR" {
				wires[wire] = input1 ^ input2
			}
		}
	}

	return wires
}

// Part 1
func combineZValues(wires map[string]int) int {
	values := make(map[int]int)
	maxZ := 0

	for key, value := range wires {
		if !strings.Contains(key, "z") {
			continue
		}
		
		zCount, _ := strconv.Atoi(key[1:])
		if zCount > maxZ {
			maxZ = zCount
		}

		values[zCount] = value
	}

	bits := make([]int, maxZ + 1)
	for key, value := range values {
		bits[key] = value
	}

	// Convert to decimal
	decimal := 0
	for i, bit := range bits {
		decimal += bit << i
	}
	fmt.Println(decimal)

	return 0
}

func part1() {
	wires, logicGates := readFile()
	allWires := calculateAvailableValues(wires, logicGates)
	combineZValues(allWires)
}

// Part 2
func detectWrongWires(wires map[string]int, logicGates map[string]LogicGate) {
	maxZ := 0

	for key := range wires {
		if !strings.Contains(key, "z") {
			continue
		}
		
		zCount, _ := strconv.Atoi(key[1:])
		if zCount > maxZ {
			maxZ = zCount
		}
	}

	var wrongWires []string
	for wire ,logicGate := range logicGates {
		if wire[0] == 'z' && logicGate.operation != "XOR" && wire != "z" + strconv.Itoa(maxZ) {
			wrongWires = append(wrongWires, wire)
		}

    if (logicGate.operation == "XOR" && 
				wire[0] != 'z' && wire[0] != 'x' && wire[0] != 'y' &&
				logicGate.input1[0] != 'z' && logicGate.input1[0] != 'x' && logicGate.input1[0] != 'y' &&
				logicGate.input2[0] != 'z' && logicGate.input2[0] != 'x' && logicGate.input2[0] != 'y') {
				wrongWires = append(wrongWires, wire)
			}

		if logicGate.operation == "AND" && logicGate.input1 != "x00" && logicGate.input2 != "x00" {
			for _, sublogicGate := range logicGates {
				if sublogicGate.operation != "OR" && (wire == sublogicGate.input1 || wire == sublogicGate.input2) {
					wrongWires = append(wrongWires, wire)
				}
			}
		}

    if logicGate.operation == "XOR" {
			for _, sublogicGate := range logicGates {
				if sublogicGate.operation == "OR" && (wire == sublogicGate.input1 || wire == sublogicGate.input2) {
					wrongWires = append(wrongWires, wire)
				}
			}
		}
	}

	// Somehow there are duplicated wires in the answer, i'm just removing them manually when sending the solution
	sort.Strings(wrongWires)
	fmt.Println(strings.Join(wrongWires, ","))
}

func part2() {
	wires, logicGates := readFile()
	allWires := calculateAvailableValues(wires, logicGates)
	detectWrongWires(allWires, logicGates)
}

func main() {
	part1()
	part2()
}