package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"math"
	// "strconv"
)

func read() (*bufio.Scanner, *os.File) {
	readFile, err := os.Open("Day04/4.txt")

	if err != nil {
		fmt.Print(err)
	}

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	return fileScanner, readFile
}



func solve(line string) int {
	line = line[strings.Index(line, ":") + 2:]
	// fmt.Println(line)

	tokenizeString := func(str, delimiters string) []string {
		return strings.FieldsFunc(str, func(r rune) bool {
			return strings.ContainsRune(delimiters, r)
		})
	}

	isInArr := func(val string, arr []string) bool {
		for _, el := range arr {
			if el == val {
				return true
			}
		}
		return false
	}

	toBeChecked := tokenizeString(line, "|")
	tokens := tokenizeString(toBeChecked[1], " ")
	toCheck := tokenizeString(toBeChecked[0], " ")

	var counter int = 0
	for _, val := range tokens {
		if isInArr(val, toCheck) {
			counter++
		}
	}

	return int(math.Pow(2, float64(counter - 1)))
}

func solver_part1() {
	fileScanner, readFile := read()

	var totalPoints int = 0
	for fileScanner.Scan() {
		totalPoints += solve(fileScanner.Text())
	}

	fmt.Println(totalPoints)

	readFile.Close()
}

func main() {
	// solver_part1()
}