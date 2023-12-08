package main

import (
	"bufio"
	"fmt"
	"os"
)

func read() (*bufio.Scanner, *os.File) {
	readFile, err := os.Open("Day08/Day08_part1/8_1.txt")

	if err != nil {
		fmt.Print(err)
	}

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	return fileScanner, readFile
}

func fetchData(line string) (key, left, right string) {
	return line[0 : 3], line[7 : 10], line[12 : 15]
}

func solver() {
	fileScanner, readFile := read()
	
	defer readFile.Close()

	var (
		instructions string
		Map = make(map[string]Node)
		lineIndex = 0
	)

	for fileScanner.Scan() {
		lineIndex++
		line := fileScanner.Text()

		if lineIndex == 1 {
			instructions = line
		}
		
		if lineIndex > 2 {
			key, left, right := fetchData(line)
			
			node := Node{
				Left: left,
				Right: right,
			}

			Map[key] = node
		}
	}

	fmt.Println(countSteps(Map, instructions))
}

type Node struct {
	Left string
	Right string
}

func countSteps(Map map[string]Node, instructions string) int {
	currentNodeLabel := "AAA"
	steps := 0

	for {
		for _, instruction := range instructions {
			if instruction == 'L' {
				currentNodeLabel = Map[currentNodeLabel].Left
			} else {
				currentNodeLabel = Map[currentNodeLabel].Right
			}

			steps++

			if currentNodeLabel == "ZZZ" {
				return steps
			}
		}
	}
}

func main() {
	solver()
}