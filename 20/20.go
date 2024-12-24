package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readFile() [][]string {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var lines [][]string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, strings.Split(line, ""))
	}

	return lines
}

func findPositions(lines [][]string, element string) []int {
	for i, line := range lines {
		for j, char := range line {
			if char == element {
				return []int{i, j}
			}
		}
	}
	panic("Element not found")
}

func arrayToString(array []int) string {
	return fmt.Sprintf("%d,%d", array[0], array[1])
}

func calculateDistances(lines [][]string) map[string]int {
	distance := make(map[string]int)
	to := findPositions(lines, "S")
	current := findPositions(lines, "E")

	accumulatedDistance := 0
	for {
		distance[arrayToString(current)] = accumulatedDistance

		if current[0] == to[0] && current[1] == to[1] {
			break
		}

		accumulatedDistance++

		for _, direction := range [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} {
			next := []int{current[0] + direction[0], current[1] + direction[1]}
			if next[0] < 0 || next[0] >= len(lines) || next[1] < 0 || next[1] >= len(lines[0]) {
				continue
			}

			if _, ok := distance[arrayToString(next)]; ok {
				continue
			}

			if lines[next[0]][next[1]] == "#" {
				continue
			}

			current = next
			break
		}
	}

	return distance
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func distance(a []int, b []int) int {
	return abs(a[0] - b[0]) + abs(a[1] - b[1])
}

func getValidCheats(lines [][]string, position []int, distances map[string]int, visited map[string]bool, distanceBetweenPoints int, currentDistance int) []int {
	var validCheats []int

	for key := range distances {
		cheat := strings.Split(key, ",")
		coordX, _ := strconv.Atoi(cheat[0])
		coordY, _ := strconv.Atoi(cheat[1])

		next := []int{coordX, coordY}

		if next[0] < 0 || next[0] >= len(lines) || next[1] < 0 || next[1] >= len(lines[0]) {
			continue
		}

		if visited[arrayToString(next)] {
			continue
		}

		calculatedDistance := distance(next, position)
		if calculatedDistance > distanceBetweenPoints {
			continue
		}

		newDistance := distances[arrayToString(next)]
		validCheats = append(validCheats, newDistance + currentDistance + calculatedDistance)
	}

	return validCheats
}

func race(lines [][]string, distances map[string]int, maxCheatDistance int) []int{
	var distancesToGoal []int
	currentPosition := findPositions(lines, "S")
	goal := findPositions(lines, "E")
	visited := make(map[string]bool)
	currentDistance := 0

	for {
		visited[arrayToString(currentPosition)] = true

		if currentPosition[0] == goal[0] && currentPosition[1] == goal[1] {
			break
		}

		validCheats := getValidCheats(lines, currentPosition, distances, visited, maxCheatDistance, currentDistance)
		for _, cheat := range validCheats {
			distancesToGoal = append(distancesToGoal, cheat)
		}

		for _, direction := range [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} {
			next := []int{currentPosition[0] + direction[0], currentPosition[1] + direction[1]}
			if next[0] < 0 || next[0] >= len(lines) || next[1] < 0 || next[1] >= len(lines[0]) {
				continue
			}

			if lines[next[0]][next[1]] == "#" {
				continue
			}

			if visited[arrayToString(next)] {
				continue
			}

			currentPosition = next
			break
		}
		currentDistance++
	}

	return distancesToGoal
}

func countCheatsThatSave100Picoseconds(distancesToGoal []int, totalDistance int) int {
	total := 0

	for _, distance := range distancesToGoal {
		savedDistance := totalDistance - distance
		if savedDistance >= 100 {
			total++
		}
	}

	return total
}

func runPart(cheatDistance int){
	lines := readFile()
	distances := calculateDistances(lines)
	distancesToGoal := race(lines, distances, cheatDistance)
	count := countCheatsThatSave100Picoseconds(distancesToGoal, distances[arrayToString(findPositions(lines, "S"))])
	fmt.Println(count)
}

func main(){
	runPart(2)
	runPart(20)
}