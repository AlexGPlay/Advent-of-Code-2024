package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

// Utils
func readFile() ([][]string, []string, []int){
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	content, _ := io.ReadAll(file)
	separation := strings.Split(string(content), "\n\n")

	var position []int

	var board [][]string
	for i, line := range strings.Split(separation[0], "\n"){
		lineContent := strings.Split(line, "")
		for j, element := range lineContent{
			if element == "@" {
				position = []int{i, j}
			}
		}

		board = append(board, lineContent)
	}

	var movements []string
	for _, line := range strings.Split(separation[1], "\n"){
		lineMovements := strings.Split(line, "")
		movements = append(movements, lineMovements...)
	}

	return board, movements, position
}

func getMovement(movement string) []int {
	if movement == "^" {
		return []int{-1, 0}
	}
	if movement == "v" {
		return []int{1, 0}
	}
	if movement == "<" {
		return []int{0, -1}
	}
	if movement == ">" {
		return []int{0, 1}
	}
	return []int{0, 0}
}

func sumBoxes(board [][]string) int {
	total := 0

	for i, line := range board{
		for j, element := range line{
			if element == "O" || element == "[" {
				total += 100 * i + j
			}
		}
	}

	return total
}

// Part 1
func move(board [][]string, position []int, movement []int) bool {
	element := board[position[0]][position[1]]
	if element == "." {
		return true
	}
	if element == "#" {
		return false
	}

	canMove := move(board, []int{position[0] + movement[0], position[1] + movement[1]}, movement)
	if canMove {
		board[position[0]][position[1]] = "."
		board[position[0] + movement[0]][position[1] + movement[1]] = element
		return true
	}
	return false
}


func part1(){
	board, movements, position := readFile()
	for _, movement := range movements{
		intMovement := getMovement(movement)
		moved := move(board, position, intMovement)
		if moved {
			position = []int{position[0] + intMovement[0], position[1] + intMovement[1]}
		}
	}
	fmt.Println(sumBoxes(board))
}

// Part 2
func prepareBoard(board [][]string) ([][]string, []int) {
	var bigBoard [][]string
	var position []int

	for i, line := range board{
		var boardLine []string
		j := 0
		for _, element := range line{
			if element == "#" {
				boardLine = append(boardLine, "#")
				boardLine = append(boardLine, "#")
			}
			if element == "." {
				boardLine = append(boardLine, ".")
				boardLine = append(boardLine, ".")
			}
			if element == "O" {
				boardLine = append(boardLine, "[")
				boardLine = append(boardLine, "]")
			}
			if element == "@" {
				boardLine = append(boardLine, "@")
				boardLine = append(boardLine, ".")
				position = []int{i, j}
			}

			j += 2
		}

		bigBoard = append(bigBoard, boardLine)
	}

	return bigBoard, position
}

func canMove(board [][]string, position []int, movement []int) bool {
	element := board[position[0]][position[1]]
	if element == "." {
		return true
	}
	if element == "#" {
		return false
	}
	if element == "[" {
		if movement[0] == 1 || movement[0] == -1 {
			return canMove(board, []int{position[0] + movement[0], position[1]}, movement) && canMove(board, []int{position[0] + movement[0], position[1] + 1}, movement)
		} else {
			return canMove(board, []int{position[0], position[1] + movement[1]}, movement)
		}
	}
	if element == "]" {
		if movement[0] == 1 || movement[0] == -1 {
			return canMove(board, []int{position[0] + movement[0], position[1]}, movement) && canMove(board, []int{position[0] + movement[0], position[1] - 1}, movement)
		} else {
			return canMove(board, []int{position[0], position[1] + movement[1]}, movement)
		}
	}

	canMove := canMove(board, []int{position[0] + movement[0], position[1] + movement[1]}, movement)
	return canMove
}

func moveWithoutChecking(board [][]string, position []int, movement []int){
	element := board[position[0]][position[1]]
	if element == "." {
		return
	}
	if element == "#" {
		return
	}

	nextPosition := []int{position[0] + movement[0], position[1] + movement[1]}
	nextElement := board[nextPosition[0]][nextPosition[1]]

	if nextElement == "[" {
		if movement[0] == 1 || movement[0] == -1 {
			moveWithoutChecking(board, []int{position[0] + movement[0], position[1] + 1}, movement)
			moveWithoutChecking(board, []int{position[0] + movement[0], position[1]}, movement)
		} else {
			moveWithoutChecking(board, []int{position[0], position[1] + movement[1]}, movement)
		}
	}
	if nextElement == "]" {
		if movement[0] == 1 || movement[0] == -1 {
			moveWithoutChecking(board, []int{position[0] + movement[0], position[1] - 1}, movement)
			moveWithoutChecking(board, []int{position[0] + movement[0], position[1]}, movement)
		} else {
			moveWithoutChecking(board, []int{position[0], position[1] + movement[1]}, movement)
		}
	}

	board[position[0]][position[1]] = "."
	board[nextPosition[0]][nextPosition[1]] = element
}

func part2(){
	board, movements, _ := readFile()
	bigBoard, position := prepareBoard(board)

	for _, movement := range movements{
		intMovement := getMovement(movement)
		isPossible := canMove(bigBoard, position, intMovement)
		if isPossible {
			moveWithoutChecking(bigBoard, position, intMovement)
			position = []int{position[0] + intMovement[0], position[1] + intMovement[1]}
		}
	}
	fmt.Println(sumBoxes(bigBoard))
}

func printBoard(board [][]string){
	for _, line := range board{
		fmt.Println(line)
	}
}

func main(){
	part1()
	part2()
}