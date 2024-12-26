package main

import (
	"bufio"
	"os"
	"sort"
	"strings"
)

// Utils
func readFile() [][]string {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var data [][]string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data = append(data, strings.Split(scanner.Text(), "-"))
	}

	return data
}

func createConnections(data [][]string) map[string][]string {
	connections := make(map[string][]string)

	for _, d := range data {
		connections[d[0]] = append(connections[d[0]], d[1])
		connections[d[1]] = append(connections[d[1]], d[0])
	}

	return connections
}

func includesAll(arr []string, elems []string) bool {
	found := make(map[string]bool)

	for _, v := range arr {
		for _, elem := range elems {
			if v == elem {
				found[elem] = true
			}
		}
	}

	return len(found) == len(elems)
}

func buildKey(elems []string) string{
	// Sort alphabetically
	sort.Strings(elems)
	return strings.Join(elems, ",")
}

func findConnectionsWith(connections map[string][]string, elems []string) map[string]bool {
	newConnections := make(map[string]bool)

	for k, v := range connections {
		if includesAll(v, elems) {
			allElems := append([]string{k}, elems...)
			newConnections[buildKey(allElems)] = true
		}
	}

	return newConnections
}

// Part 1
func makeTripleConnections(connections map[string][]string) map[string]bool {
	newConnections := make(map[string]bool)

	for k, v := range connections {
		elem1 := k
		for _, elem2 := range v {
			triConnections := findConnectionsWith(connections, []string{elem1, elem2})
			for k := range triConnections {
				newConnections[k] = true
			}
		}
	}

	return newConnections
}

func countConnectionsThatStartsWithT(connections map[string]bool) int {
	total := 0

	for k := range connections {
		elems := strings.Split(k, ",")
		for _, elem := range elems {
			if elem[0] == 't' {
				total++
				break
			}
		}
	}

	return total
}

func part1(){
	data := readFile()
	connections := createConnections(data)
	tripleConnections := makeTripleConnections(connections)	
	total := countConnectionsThatStartsWithT(tripleConnections)
	println(total)
}

// Part 2
func discoverBiggestConnection(connections map[string]int) string {
	biggest := ""
	biggestSize := 0

	for k, v := range connections {
		if v > biggestSize {
			biggest = k
			biggestSize = v
		}
	}

	return biggest
}

func exploreConnections(connections map[string][]string, elems []string, visited map[string]bool) map[string]int {
	newConnections := make(map[string]int)

	for k, v := range connections {
		if visited[k] {
			continue
		}
		visited[k] = true
		if includesAll(v, elems) {
			allElems := append([]string{k}, elems...)
			newConnections[buildKey(allElems)] = len(allElems)
			moreConnections := exploreConnections(connections, allElems, visited)
			for k, v := range moreConnections {
				newConnections[k] = v
			}
		}
	}

	return newConnections
}

func makeFullConnections(connections map[string][]string) map[string]int {
	newConnections := make(map[string]int)
	
	for k := range connections {
		visited := make(map[string]bool)
		visited[k] = true
		connections := exploreConnections(connections, []string{k}, visited)
		for k, v := range connections {
			newConnections[k] = v
		}
	}

	return newConnections
}

func part2(){
	data := readFile()
	connections := createConnections(data)
	fullConnections := makeFullConnections(connections)
	biggestConnection := discoverBiggestConnection(fullConnections)
	println(biggestConnection)
}

func main(){
	part1()
	part2()
}