package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func read() (*bufio.Scanner, *os.File) {
	readFile, err := os.Open("Day08/Day08_part2/8_2.txt")

	if err != nil {
		fmt.Print(err)
	}

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	return fileScanner, readFile
}

func fetchData(line string) (key, left, right string) {
	return line[0:3], line[7:10], line[12:15]
}

type Node struct {
	Left  string
	Right string
}

func solver() {
	fileScanner, readFile := read()
	defer readFile.Close()

	var (
		instructions string
		Map          = make(map[string]Node)
		lineIndex    = 0
	)

	for fileScanner.Scan() {
		lineIndex++
		line := fileScanner.Text()

		if lineIndex == 1 {
			instructions = line
		}

		if lineIndex > 2 {
			key, left, right := fetchData(line)

			node := Node {
				Left:  left,
				Right: right,
			}

			Map[key] = node
		}
	}

	fmt.Println(countStepsToEndZ(Map, instructions))
}

func getPathStepFrom(startNode string, Map map[string]Node, instructions string) int {
	currentNodeLabel := startNode
	steps := 1

	for {
		for _, instruction := range instructions {
			if instruction == 'L' {
				currentNodeLabel = Map[currentNodeLabel].Left
			} else {
				currentNodeLabel = Map[currentNodeLabel].Right
			}
	
			steps++

			if strings.HasSuffix(currentNodeLabel, "Z") {
				return steps - 1
			}
		}
	}
}

func countStepsToEndZ(Map map[string]Node, instructions string) int {
	var startingNodes []string
	for node := range Map {
		if strings.HasSuffix(node, "A") {
			startingNodes = append(startingNodes, node)
		}
	}

	var results []int
	for _, startingNode := range startingNodes {
		results = append(results, getPathStepFrom(startingNode, Map, instructions))
	}

	return findLCM(results...)
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a % b
	}
	return a
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

func findLCM(arr ...int) int {
	if len(arr) == 0 {
		return 0 
	}

	result := arr[0]
	for i := 1; i < len(arr); i++ {
		result = lcm(result, arr[i])
	}

	return result
}

func main() {
	solver()
}