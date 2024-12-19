package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
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
		lines = append(lines, strings.Split(scanner.Text(), ""))
	}
	return lines
}

func encodePosition(i int, j int) string {
	return fmt.Sprintf("%d,%d", i, j)
}

func cloneVisited(visited map[string]bool) map[string]bool {
	clone := make(map[string]bool)
	for key, value := range visited {
		clone[key] = value
	}
	return clone
}

func findNode(idx int, nodeMap map[string]int) string {
	for key, value := range nodeMap {
		if value == idx {
			return key
		}
	}
	return ""
}

func findDirectionMovement(from []int, to []int) string {
	if from[0] == to[0] {
		return "horizontal"
	}
	return "vertical"
}

type QueueElement struct {
	position []int
	weight int
	visited map[string]bool
	direction string
}

func findNeighbours(lines [][]string, weight int, position []int, visited map[string] bool, direction string) []QueueElement {
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

		newDirection := "vertical"
		if movement[1] != 0 {
			newDirection = "horizontal"
		}

		movementCost := weight + 1
		if newDirection != direction {
			movementCost += 1000
		}

		neighbours = append(neighbours, QueueElement{newPosition, movementCost, visited, newDirection})
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

func dijkstra(lines [][]string, from []int, to []int) (int, int) {
	minCost := math.MaxInt
	weights := make(map[string]int)
	minPaths := make(map[string]bool)
	queue := []QueueElement{{from, 0, make(map[string]bool), "horizontal"}}

	for len(queue) > 0 {
		element := queue[0]
		queue = queue[1:]

		if element.weight > minCost {
			continue
		}

		elementWeightKey := encodePosition(element.position[0], element.position[1]) + element.direction
		if _, ok := weights[elementWeightKey]; ok && weights[elementWeightKey] < element.weight {
			continue
		}
		weights[elementWeightKey] = element.weight

		for _, neighbour := range findNeighbours(lines, element.weight, element.position, element.visited, element.direction) {
			isEnd := neighbour.position[0] == to[0] && neighbour.position[1] == to[1]
			if isEnd {
				if element.weight < minCost {
					minCost = element.weight
					minPaths = cloneVisited(element.visited)
				} else if element.weight == minCost {
					for key, value := range element.visited {
						minPaths[key] = value
					}
				}
			} else {
				cloneVisited := cloneVisited(element.visited)
				cloneVisited[encodePosition(neighbour.position[0], neighbour.position[1])] = true
				if element.weight < minCost {
					queue = append(queue, QueueElement{neighbour.position, neighbour.weight, cloneVisited, neighbour.direction})
				}
			}
		}

	}

	return minCost, len(minPaths) + 2
}

func findElement(lines [][]string, element string) []int {
	for i, line := range lines {
		for j, char := range line {
			if char == element {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

func main(){
	lines := readFile()
	from := findElement(lines, "S")
	to := findElement(lines, "E")
	// This takes forever because of part 2 and at this point i'm too tired to optimize it
	minCost, paths := dijkstra(lines, from, to)
	fmt.Println(minCost)
	fmt.Println(paths)
}