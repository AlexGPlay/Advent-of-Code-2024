package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func readFile() ([]string, []string) {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	contents, _ := io.ReadAll(file)
	sections := strings.Split(string(contents), "\n\n")

	availableTowels := strings.Split(sections[0], ", ")
	designs := strings.Split(sections[1], "\n")
	
	return availableTowels, designs
}

func canBeMade(towels []string, design string, cache map[string]int) int {
	if value, ok := cache[design]; ok {
		return value
	}

	for _, towel := range towels {
		if strings.HasPrefix(design, towel) {
			pendingDesign := strings.TrimPrefix(design, towel)
			if pendingDesign == "" {
				if value, ok := cache[design]; ok {
					cache[design] = value + 1
				} else {
					cache[design] = 1
				}
			}
			count := canBeMade(towels, pendingDesign, cache)
			if value, ok := cache[design]; ok {
				cache[design] = value + count
			} else {
				cache[design] = count
			}
		}
	}

	if value, ok := cache[design]; ok {
		return value
	}
	
	return 0
}

func part1() {
	towels, designs := readFile()
	cache := make(map[string]int)
	count := 0
	for _, design := range designs {
		if canBeMade(towels, design, cache) > 0 {
			count++
		}
	}
	fmt.Println(count)
}

func part2() {
	towels, designs := readFile()
	cache := make(map[string]int)
	count := 0
	for _, design := range designs {
		designCount := canBeMade(towels, design, cache)
		count += designCount
	}
	fmt.Println(count)
}

func main() {
	part1()
	part2()
}