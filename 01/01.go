package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func part1() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()

	var column1 []int
	var column2 []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		var value1, value2 int

		fmt.Sscanf(line, "%d %d", &value1, &value2)
		column1 = append(column1, value1)
		column2 = append(column2, value2)
	}

	sort.Ints(column1)
	sort.Ints(column2)

	total := 0
	for i, value1 := range column1 {
		result := abs(value1 - column2[i])
		total += result
	}

	fmt.Println(total)
}

func part2(){
	var column1 []int
	columnRepetitions:= make(map[int]int)

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		var value1, value2 int

		fmt.Sscanf(line, "%d %d", &value1, &value2)
		column1 = append(column1, value1)

		columnRepetitions[value2]++
	}

	total := 0
	for _, value1 := range column1 {
		repetitions, ok := columnRepetitions[value1]
		if ok {
			total += repetitions * value1
		}
	}
	fmt.Println(total)
}

func main(){
	part1()
	part2()
}