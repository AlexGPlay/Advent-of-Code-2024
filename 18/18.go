package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Utils
func readFile() [][]int {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var grid [][]int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		numbersAsString := strings.Split(line, ",")
		var pair []int
		for _, numberAsString := range numbersAsString {
			number, _ := strconv.Atoi(numberAsString)
			pair = append(pair, number)
		}

		grid = append(grid, pair)
	}

	return grid
}

func buildGrid(size int, falledByes [][]int) [][]string{
	var grid [][]string
	for i := 0; i <= size; i++ {
		var row []string
		for j := 0; j <= size; j++ {
			row = append(row, ".")
		}
		grid = append(grid, row)
	}

	for _, pair := range falledByes {
		x, y := pair[0], pair[1]
		grid[y][x] = "#"
	}

	return grid
}

type QueueElement struct {
	position []int
	weight int
}

func encodePosition(i int, j int) string {
	return fmt.Sprintf("%d,%d", i, j)
}

func findNeighbours(lines [][]string, weight int, position []int) []QueueElement {
	var neighbours []QueueElement

	potentialMovements := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for _, movement := range potentialMovements {
		newPosition := []int{position[0] + movement[0], position[1] + movement[1]}
		if newPosition[0] < 0 || newPosition[0] >= len(lines) || newPosition[1] < 0 || newPosition[1] >= len(lines[0]) {
			continue
		}

		if lines[newPosition[0]][newPosition[1]] == "#" {
			continue
		}

		movementCost := weight + 1
		neighbours = append(neighbours, QueueElement{newPosition, movementCost})
	}


	return neighbours
}

func sortQueue(queue []QueueElement) []QueueElement {
	for i := 0; i < len(queue); i++ {
		for j := i + 1; j < len(queue); j++ {
			if queue[i].weight > queue[j].weight {
				queue[i], queue[j] = queue[j], queue[i]
			}
		}
	}
	return queue
}

func dijkstra(lines [][]string, from []int, to []int) int {
	queue := []QueueElement{{from, 0}}
	visited := make(map[string]bool)
	weights := make(map[string]int)

	for len(queue) > 0 {
		element := queue[0]
		queue = queue[1:]

		key := encodePosition(element.position[0], element.position[1])

		if _, ok := weights[key]; ok {
			if weights[key] < element.weight {
				continue
			} else {
				weights[key] = element.weight
			}
		} else {
			weights[key] = element.weight
		}

		if visited[key] {
			continue
		}

		visited[key] = true

		neighbours := findNeighbours(lines, weights[key], element.position)
		for _, neighbour := range neighbours {
			queue = append(queue, neighbour)
		}
		queue = sortQueue(queue)
	}


	return weights[encodePosition(to[0], to[1])]
}

// Part 1
func part1(size int, fallenBytes int){
	positions := readFile()
	usefulPositions := positions[:fallenBytes]
	grid := buildGrid(size, usefulPositions)
	shortestPath := dijkstra(grid, []int{0, 0}, []int{size, size})

	fmt.Println(shortestPath)
}

// Part 2
func removeElementFromGrid(grid [][]string, element []int) [][]string{
	grid[element[1]][element[0]] = "."
	return grid
}

func part2(size int) {
	positions := readFile()
	grid := buildGrid(size, positions)
	i := len(positions) - 1

	for i >= 0 {
		grid := removeElementFromGrid(grid, positions[i])
		shortestPath := dijkstra(grid, []int{0, 0}, []int{size, size})
		if shortestPath != 0 {
			fmt.Printf("%d,%d", positions[i][0], positions[i][1])
			break
		}

		i--
	}
}

func main(){
	part1(70, 1024) //part1(6, 12) => The example is a 6x6 grid and 12 fallen bytes
	part2(70) //part2(6) => The example is a 6x6 grid
}