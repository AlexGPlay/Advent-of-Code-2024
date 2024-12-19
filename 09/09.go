package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

// Utils
type Block struct {
	blockType string
	id int
	size int
}

func readInput() []int{
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var input []int
	content, _ := io.ReadAll(file)
	elems := strings.Split(string(content), "")

	for _, elem := range elems {
		num, _ := strconv.Atoi(elem)
		input = append(input, num)
	}

	return input
}

func copy(blocks []Block) []Block {
	var newBlocks []Block
	for _, block := range blocks {
		newBlock := Block{block.blockType, block.id, block.size}
		newBlocks = append(newBlocks, newBlock)
	}
	return newBlocks
}

func buildBlocks(input []int) []Block {
	var blocks []Block
	for i, value := range input {
		if i % 2 == 0 {
			blocks = append(blocks, Block{"block", i / 2, value})
		} else {
			blocks = append(blocks, Block{"gap", -1, value})
		}
	}
	return blocks
}

// Part 1
func fillGap(blocks []Block, gapIdx int) []Block {
	tmpBlocks := copy(blocks)
	neededSize := tmpBlocks[gapIdx].size

	var toReplace []Block
	for i := len(blocks) - 1; i > gapIdx; i-- {
		block := tmpBlocks[i]
		if block.blockType == "gap" {
			continue
		}

		if block.size == 0 {
			continue
		}

		substraction := math.Min(float64(block.size), float64(neededSize))
		neededSize -= int(substraction)
		tmpBlocks[i].size -= int(substraction)

		toReplace = append(toReplace, Block{"block", block.id, int(substraction)})

		if neededSize == 0 {
			break
		}
	}

	var newBlocks []Block
	newBlocks = append(newBlocks, tmpBlocks[:gapIdx]...)
	newBlocks = append(newBlocks, toReplace...)
	newBlocks = append(newBlocks, tmpBlocks[gapIdx+1:]...)

	return newBlocks
}

func part1(){
	input := readInput()
	blocks := buildBlocks(input)

	i := 0

	for {
		block := blocks[i]
		if block.blockType == "gap" {
			blocks = fillGap(blocks, i)
		}

		i++
		if i > len(blocks) - 1 {
			break
		}
	}

	total := 0
	i = 0

	for _, value := range blocks {
		for j := 0; j < value.size; j++ {
			total += value.id * i
			i++
		}
	}

	fmt.Println(total)
}

// Part 2
func fillLeftmostGap(blocks []Block, blockIdx int) ([]Block, bool) {
	tmpBlocks := copy(blocks)
	neededSize := tmpBlocks[blockIdx].size

	var newBlocks []Block

	for i := 0; i < blockIdx; i++ {
		block := tmpBlocks[i]
		if block.blockType != "gap" {
			continue
		}

		canFitBlock := block.size >= neededSize
		if !canFitBlock {
			continue
		}

		remainingSpace := block.size - neededSize

		newBlocks = append(newBlocks, tmpBlocks[:i]...)
		newBlocks = append(newBlocks, Block{"block", tmpBlocks[blockIdx].id, neededSize})

		if remainingSpace > 0 {
			newBlocks = append(newBlocks, Block{"gap", block.id, remainingSpace})
		}

		newBlocks = append(newBlocks, tmpBlocks[i+1:blockIdx]...)
		newBlocks = append(newBlocks, Block{"gap", -1, neededSize})
		newBlocks = append(newBlocks, tmpBlocks[blockIdx + 1:]...)

		break
	}

	if len(newBlocks) == 0 {
		return blocks, false
	}

	return newBlocks, true
}

func part2(){
	input := readInput()
	blocks := buildBlocks(input)

	movedMap := make(map[int]bool)

	i := len(blocks) - 1
	for {
		if i < 0 {
			break
		}

		block := blocks[i]
		if(block.blockType == "gap") {
			i--
			continue
		}

		if _, ok := movedMap[block.id]; ok {
			i--
			continue
		}

		newBlocks, moved := fillLeftmostGap(blocks, i)
		blocks = newBlocks
		if moved {
			movedMap[block.id] = true
		} else {
			i--
		}
	}

	total := 0
	i = 0

	for _, value := range blocks {
		if value.blockType == "gap" {
			i += value.size
			continue
		}

		for j := 0; j < value.size; j++ {
			total += value.id * i
			i++
		}
	}

	fmt.Println(total)
}


func main(){
	part1()
	part2()
}