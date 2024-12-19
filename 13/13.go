package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// Utils
type Machine struct {
	a []float64
	b []float64
	prize []float64
}

func readFile() []Machine{
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Error reading file")
	}
	defer file.Close()

	content, _ := io.ReadAll(file)
	contentAsString := string(content)

	blocks := strings.Split(contentAsString, "\n\n")

	var machines []Machine
	for _, block := range blocks {
		lines := strings.Split(block, "\n")

		// Get the numbers with regex
		numbersAString := regexp.MustCompile(`\d+`).FindAllString(lines[0], -1)
		var numbersA []float64
		for _, numberAsString := range numbersAString {
			number, _ := strconv.Atoi(numberAsString)
			numbersA = append(numbersA, float64(number))
		}

		numbersBString := regexp.MustCompile(`\d+`).FindAllString(lines[1], -1)
		var numbersB []float64
		for _, numberBString := range numbersBString {
			number, _ := strconv.Atoi(numberBString)
			numbersB = append(numbersB, float64(number))
		}

		numbersPricesString := regexp.MustCompile(`\d+`).FindAllString(lines[2], -1)
		var numbersPrices []float64
		for _, numberPriceString := range numbersPricesString {
			number, _ := strconv.Atoi(numberPriceString)
			numbersPrices = append(numbersPrices, float64(number))
		}

		machines = append(machines, Machine{a: numbersA, b: numbersB, prize: numbersPrices})
	}

	return machines
}

func countDecimals(f float64) int {
	str := fmt.Sprintf("%.15g", f)

	if idx := strings.Index(str, "."); idx != -1 {
		return len(str) - idx - 1
	}

	return 0
}

func solveMachine(machine Machine, maxTries float64) float64{
	ax, ay := machine.a[0], machine.a[1]
	bx, by := machine.b[0], machine.b[1]
	px, py := machine.prize[0], machine.prize[1]	

	bSolution := (py*ax-ay*px)/(by*ax-ay*bx)
	aSolution := (px-bx*bSolution)/ax

	if countDecimals(aSolution) > 0 || countDecimals(bSolution) > 0 {
		return 0
	}

	if maxTries > 0 {
		if aSolution > maxTries || bSolution > maxTries {
			return 0
		}
	}

	tokens := aSolution * 3 + bSolution
	return tokens
}

// Part 1
func part1(){
	machines := readFile()

	total := 0.0
	for _, machine := range machines {
		total += solveMachine(machine, 100)
	}
	fmt.Println(total)
}

// Part 2
func part2(){
	machines := readFile()

	total := 0.0
	for _, machine := range machines {
		machine.prize[0] += 10000000000000
		machine.prize[1] += 10000000000000
		total += solveMachine(machine, -1)
	}
	fmt.Println(int64(total))
}

func main(){
	part1()
	part2()	
}