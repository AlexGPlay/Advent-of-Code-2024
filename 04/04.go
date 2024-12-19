package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// Utils
func parseFile() [][]string {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()

	var parsedContent [][]string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "")
		parsedContent = append(parsedContent, line)
	}

	return parsedContent
}

// Part 1
func countHorizontalMatches(lines [][]string) int{
	total := 0

	for _, line := range lines {
		lineAsString := strings.Join(line, "")
		matches := strings.Count(lineAsString, "XMAS") + strings.Count(lineAsString, "SAMX")
		total += matches
	}

	return total
}

func countVerticalMatches(lines [][]string) int{
	total := 0

	for i := 0; i < len(lines[0]); i++ {
		var column []string
		for j := 0; j < len(lines); j++ {
			column = append(column, lines[j][i])
		}

		columnAsString := strings.Join(column, "")
		matches := strings.Count(columnAsString, "XMAS") + strings.Count(columnAsString, "SAMX")
		total += matches
	}

	return total
}

func countDiaognalMatchesLeft(lines [][]string) int {
	total := 0

	for i:=0; i < len(lines); i++ {
		for j:=0; j < len(lines[0]); j++ {
			if(i + 3 < len(lines) && j + 3 < len(lines[0])){
				if(lines[i][j] == "X" && lines[i+1][j+1] == "M" && lines[i+2][j+2] == "A" && lines[i+3][j+3] == "S"){
					total++
				}
			}

			if(i + 3 < len(lines) && j + 3 < len(lines[0])){
				if(lines[i][j] == "S" && lines[i+1][j+1] == "A" && lines[i+2][j+2] == "M" && lines[i+3][j+3] == "X"){
					total++
				}
			}
		}
	}

	return total
}

func countDiaognalMatchesRight(lines [][]string) int {
	total := 0

	for i:=0; i < len(lines); i++ {
		for j:=0; j < len(lines[0]); j++ {
			if(i + 3 < len(lines) && j - 3 >= 0){
				if(lines[i][j] == "X" && lines[i+1][j-1] == "M" && lines[i+2][j-2] == "A" && lines[i+3][j-3] == "S"){
					total++
				}
			}

			if(i + 3 < len(lines) && j - 3 >= 0){
				if(lines[i][j] == "S" && lines[i+1][j-1] == "A" && lines[i+2][j-2] == "M" && lines[i+3][j-3] == "X"){
					total++
				}
			}
		}
	}

	return total
}

func part1(){
	fileContent := parseFile()
	total := countHorizontalMatches(fileContent) + countVerticalMatches(fileContent) + countDiaognalMatchesLeft(fileContent) + countDiaognalMatchesRight(fileContent)
	fmt.Println(total)
}

// Part 2
func countXmasMatches(lines [][]string) int{
	total := 0

	for i:=1; i < len(lines) - 1; i++ {
		for j:=1; j < len(lines[0]) - 1; j++ {
			if(lines[i][j] == "A"){
				diagonal1 := lines[i-1][j-1] + lines[i][j] + lines[i+1][j+1]
				isValidDiagonal1 := diagonal1 == "MAS" || diagonal1 == "SAM"

				diaognal2 := lines[i-1][j+1] + lines[i][j] + lines[i+1][j-1]
				isValidDiagonal2 := diaognal2 == "MAS" || diaognal2 == "SAM"

				if(isValidDiagonal1 && isValidDiagonal2){
					total++
				}
			}
		}
	}

	return total
}

func part2(){
	fileContent := parseFile()
	total := countXmasMatches(fileContent)
	fmt.Println(total)
}

func main(){
	part1()
	part2()
}