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
func findIndex(line []string) int{
	for i, v := range line {
		if v == "^" {
			return i
		}
	}

	return -1
}

func parseFile() ([][]string, []int) {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var problemMap [][]string
	var initialPosition []int

	i := 0
	for scanner.Scan(){
		line := strings.Split(scanner.Text(), "")
		problemMap = append(problemMap, line)
		idx := findIndex(line)
		if idx != -1 {
			initialPosition = append(initialPosition, i, idx)
		}

		i++
	}

	return problemMap, initialPosition
}

func rotateRight(directionVector []int) []int{
	if directionVector[0] == 0 && directionVector[1] == 1 {
		return []int{1, 0}
	}
	if directionVector[0] == 1 && directionVector[1] == 0 {
		return []int{0, -1}
	}
	if directionVector[0] == 0 && directionVector[1] == -1 {
		return []int{-1, 0}
	}
	if directionVector[0] == -1 && directionVector[1] == 0 {
		return []int{0, 1}
	}
	// Should never happen
	return []int{0, 0}
}

func isOutOfBounds(problemMap [][]string, position []int) bool {
	return position[0] < 0 || position[0] >= len(problemMap) || position[1] < 0 || position[1] >= len(problemMap[0])
}

func move(problemMap [][]string, position []int, directionVector []int) ([]int, []int) {
	potentialNewPosition := []int{position[0] + directionVector[0], position[1] + directionVector[1]}

	if isOutOfBounds(problemMap, potentialNewPosition) {
		return potentialNewPosition, directionVector
	}

	if problemMap[potentialNewPosition[0]][potentialNewPosition[1]] == "#" {
		return move(problemMap, position, rotateRight(directionVector))
	}

	return potentialNewPosition, directionVector
}

// Part 1
func makeKey(position []int) string{
	return strconv.Itoa(position[0]) + "," + strconv.Itoa(position[1])
}

func part1() map[string]bool{
	problemMap, position := parseFile()

	currentDirection := []int{-1, 0}

	visitedPositions := make(map[string]bool)
	initialPositionKey := makeKey(position)
	visitedPositions[initialPositionKey] = true

	for {
		position, currentDirection = move(problemMap, position, currentDirection)
		if isOutOfBounds(problemMap, position) {
			break
		}

		positionKey := makeKey(position)
		visitedPositions[positionKey] = true
	}

	fmt.Println(len(visitedPositions))

	return visitedPositions
}

// Part 2
func makeKeyWithDirection(position []int, direction []int) string{
	return strconv.Itoa(position[0]) + "," + strconv.Itoa(position[1]) + "-" + strconv.Itoa(direction[0]) + "," + strconv.Itoa(direction[1])
}

func makeNewMap(problemMap [][]string, position []int) [][]string{
	newMap := make([][]string, len(problemMap))
	for i := range newMap {
		newMap[i] = make([]string, len(problemMap[i]))
		copy(newMap[i], problemMap[i])
	}

	newMap[position[0]][position[1]] = "#"

	return newMap
}

func isMapWithLoop(problemMap [][]string, position []int) bool{
	currentDirection := []int{-1, 0}

	visitedPositions := make(map[string]bool)
	initialPositionKey := makeKeyWithDirection(position, currentDirection)
	visitedPositions[initialPositionKey] = true

	for {
		position, currentDirection = move(problemMap, position, currentDirection)
		if isOutOfBounds(problemMap, position) {
			return false
		}

		positionKey := makeKeyWithDirection(position, currentDirection)
		if visitedPositions[positionKey] {
			return true
		}

		visitedPositions[positionKey] = true
	}
}

func part2(){
	problemMap, position := parseFile()
	candidates := part1()

	total := 0
	for key := range candidates {
		candidate := strings.Split(key, ",")
		firstPosition, _ := strconv.Atoi(candidate[0])
		secondPosition, _ := strconv.Atoi(candidate[1])

		newMap := makeNewMap(problemMap, []int{firstPosition, secondPosition})
		if isMapWithLoop(newMap, position) {
			total++
		}
	}

	fmt.Println(total)
}

// Main
func main() {
	part2() // Part 1 is also executed for part 2
}