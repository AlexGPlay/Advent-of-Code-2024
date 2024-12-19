package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// Utils
type Line struct {
	result  int
	numbers []int
}

func parseFile() []Line {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []Line

	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.Split(line, ": ")
		result, _ := strconv.Atoi(parts[0])

		numbersAsString := strings.Split(parts[1], " ")
		var numbers []int
		for _, numberAsString := range numbersAsString {
			number, _ := strconv.Atoi(numberAsString)
			numbers = append(numbers, number)
		}

		lines = append(lines, Line{result, numbers})
	}

	return lines
}

// Part 1
func hasValidCombination(numbers []int, result int) bool{
	if len(numbers) == 1 {
		return numbers[0] == result
	}

	sumResult := numbers[0] + numbers[1]
	newNumbers := append([]int{sumResult}, numbers[2:]...)
	if hasValidCombination(newNumbers, result) {
		return true
	}

	multResult := numbers[0] * numbers[1]
	newNumbers = append([]int{multResult}, numbers[2:]...)
	if hasValidCombination(newNumbers, result) {
		return true
	}

	return false
}

func part1() {
	data := parseFile()

	total := 0
	for _, line := range data {
		if hasValidCombination(line.numbers, line.result) {
			total += line.result
		}
	}

	fmt.Println(total)
}

// Part 2
func hasValidCombinationWithConcatenation(numbers []int, result int) bool{
	if len(numbers) == 1 {
		return numbers[0] == result
	}

	sumResult := numbers[0] + numbers[1]
	newNumbers := append([]int{sumResult}, numbers[2:]...)
	if hasValidCombinationWithConcatenation(newNumbers, result) {
		return true
	}

	multResult := numbers[0] * numbers[1]
	newNumbers = append([]int{multResult}, numbers[2:]...)
	if hasValidCombinationWithConcatenation(newNumbers, result) {
		return true
	}

	concatenationResult, _ := strconv.Atoi(strconv.Itoa(numbers[0]) + strconv.Itoa(numbers[1]))
	newNumbers = append([]int{concatenationResult}, numbers[2:]...)
	if hasValidCombinationWithConcatenation(newNumbers, result) {
		return true
	}

	return false
}

func part2() {
	data := parseFile()

	total := 0
	for _, line := range data {
		if hasValidCombinationWithConcatenation(line.numbers, line.result) {
			total += line.result
		}
	}

	fmt.Println(total)
}

// Main
func main() {
	part1()
	part2()
}