package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

// Utils
type Ordering struct {
	from int
	to int
}

func parseFile() ([]Ordering, [][]int) {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()

	content, _ := io.ReadAll(file)
	contentAsString := string(content)

	data := strings.Split(contentAsString, "\n")

	var orderings []Ordering
	var updates [][]int

	var parseMode = 0
	for _, line := range data {
		if strings.TrimSpace(line) == "" {
			parseMode = 1
			continue
		}

		if parseMode == 0 {
			values := strings.Split(strings.TrimSpace(line), "|")

			from, _ := strconv.Atoi(values[0])
			to, _ := strconv.Atoi(values[1])
	
	
			orderings = append(orderings, Ordering{from: from, to: to})
		}

		if parseMode == 1 {
			values := strings.Split(strings.TrimSpace(line), ",")
			var update []int
			for _, value := range values {
				val, _ := strconv.Atoi(value)
				update = append(update, val)
			}
			updates = append(updates, update)
		}

	}

	return orderings, updates
}

func createOrderMap(orderings []Ordering) map[int][]int {
	orderMap := make(map[int][]int)
	for _, ordering := range orderings {
		if _, ok := orderMap[ordering.from]; ok {
			orderMap[ordering.from] = append(orderMap[ordering.from], ordering.to)
		} else {
			orderMap[ordering.from] = []int{ordering.to}
		}
	}

	return orderMap
}

func contains(numbers []int, value int) bool {
	for _, number := range numbers {
		if number == value {
			return true
		}
	}

	return false
}

func isValidOrdering(section []int, numbers []int) bool {
	for _, value := range section {
		if contains(numbers, value) {
			return false
		}
	}

	return true
}

func updateScore(orderMap map[int][]int, updates []int) int {
	for i, value := range updates {
		rules, ok := orderMap[value]
		if !ok {
			continue
		}

		sliceToCheck := updates[:i]
		if !isValidOrdering(sliceToCheck, rules) {
			return 0
		}

	}

	middleValue := updates[len(updates) / 2]
	return middleValue
}

// Part 1
func part1(){
	orderings, updates := parseFile()
	orderMap := createOrderMap(orderings)

	total := 0
	for _, update := range updates {
		score := updateScore(orderMap, update)
		total += score
	}

	fmt.Println(total)
}

// Part 2
func sort(orderMap map[int][]int, updates []int) []int{
	n := len(updates)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			orderMapValues, ok := orderMap[updates[j + 1]]
			if !ok {
				continue
			}

			shouldGoAfter := contains(orderMapValues, updates[j])

			if shouldGoAfter {
				updates[j], updates[j+1] = updates[j+1], updates[j]
			}
		}
	}

	return updates
}

func part2(){
	orderings, updates := parseFile()
	orderMap := createOrderMap(orderings)

	total := 0
	for _, update := range updates {
		score := updateScore(orderMap, update)
		if score == 0 {
			sorted := sort(orderMap, update)
			total += updateScore(orderMap, sorted)
		}
	}

	fmt.Println(total)
}

// Main
func main(){
	part1()
	part2()
}