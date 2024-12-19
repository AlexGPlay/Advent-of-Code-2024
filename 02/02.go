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
func parseFile() [][]int {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()

	var rows [][]int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var row []int
		
		line := strings.Split(scanner.Text(), " ")
		for _, value := range line {
			value, _ := strconv.Atoi(value)
			row = append(row, value)
		}

		rows = append(rows, row)
	}	

	return rows
}

func abs (x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func isIncreasing(row []int) bool {
	for i := 0; i < len(row)-1; i++ {
		if row[i] > row[i+1] {
			return false
		}
	}
	return true
}

func isDecreasing(row []int) bool {
	for i := 0; i < len(row)-1; i++ {
		if row[i] < row[i+1] {
			return false
		}
	}
	return true
}

func isSafe(row []int) bool {
	for i := 0; i < len(row)-1; i++ {
		value := row[i]
		nextValue := row[i+1]

		diff := abs(nextValue - value)
		if diff < 1 || diff > 3 {
			return false
		}
	}

	return true
}

func isSafeRow(row []int) bool {
	if isIncreasing(row) && isSafe(row) {
		return true
	}
	if isDecreasing(row) && isSafe(row) {
		return true
	}
	return false
}

// Part 1
func part1(){
	rows := parseFile()
	
	// Process it
	total := 0
	for _, row := range rows {
		if isSafeRow(row){
			total += 1
		}
	}

	fmt.Println(total)
}

// Part 2
// Brute force removing 1 by 1 if it initially fails
func isSafeByRemovingOne(row []int) bool {
	for i := 0; i < len(row); i++ {
		var sliceWithoutPosition []int

		sliceWithoutPosition = append(sliceWithoutPosition, row[:i]...)
		sliceWithoutPosition = append(sliceWithoutPosition, row[i+1:]...)

		if isSafeRow(sliceWithoutPosition) {
			return true
		}
	}

	return false
}

func part2(){
	rows := parseFile()
	
	// Process it
	total := 0
	for _, row := range rows {
		if isSafeRow(row){
			total += 1
			continue
		}
		if isSafeByRemovingOne(row) {
			total += 1
			continue
		}
	}

	fmt.Println(total)
}

// Main
func main(){
	part1()
	part2()
}