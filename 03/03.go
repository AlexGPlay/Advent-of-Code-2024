package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
)

// Utils
func parseFile() string {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	return string(content)
}

// Part 1
func part1(){
	fileContent := parseFile()

	pattern := `mul\((?P<num1>\d+),(?P<num2>\d+)\)`
	re := regexp.MustCompile(pattern)
	matches := re.FindAllStringSubmatch(fileContent, -1)

	total := 0
	for _, match := range matches {
		num1, _ := strconv.Atoi(match[1])
		num2, _ := strconv.Atoi(match[2])

		total += num1 * num2
	}

	fmt.Println(total)
}

// Part 2
func part2(){
	fileContent := parseFile()

	pattern := `mul\((?P<num1>\d+),(?P<num2>\d+)\)|do\(\)|don't\(\)`
	re := regexp.MustCompile(pattern)
	matches := re.FindAllStringSubmatch(fileContent, -1)

	enabled := true
	total := 0
	for _, match := range matches {
		if(match[0] == "do()"){
			enabled = true
		} else if(match[0] == "don't()"){
			enabled = false
		} else{
			if enabled {
				num1, _ := strconv.Atoi(match[1])
				num2, _ := strconv.Atoi(match[2])

				total += num1 * num2
			}
		}
	}

	fmt.Println(total)
}

// Main
func main(){
	part1()
	part2()
}