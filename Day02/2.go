package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func read() (*bufio.Scanner, *os.File) {
	readFile, err := os.Open("Day02/2.txt")

	if err != nil {
		fmt.Print(err)
	}

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	return fileScanner, readFile
}

func getNum(str string) int {
	var num int = 0
	for _, char := range str {
		digit := int(char - '0')

		if 0 <= digit && digit <= 9 {
			num = num * 10 + digit
		}
	}

	return num
}

func isPossible_part1(line string) bool {
	colors := map[string]int{
		"green": 13,
		"red":   12,
		"blue":  14,
	}

	line = line[strings.Index(line, ":") + 2:]

	const delimiters = ",;"

	tokenizeString := func(str, delimiters string) []string {
		return strings.FieldsFunc(str, func(r rune) bool {
			return strings.ContainsRune(delimiters, r)
		})
	}

	tokens := tokenizeString(line, delimiters)

	for _, token := range tokens {
		for key := range colors {
			if getNum(token) > colors[key] {
				return false
			}				
		}
	}

	return true
}

func cubePowerPart2(line string) int {
	const limit int = -1
	colors := map[string]int{
		"green": limit,
		"red":   limit,
		"blue":  limit,
	}

	line = line[strings.Index(line, ":") + 2:]

	const delimiters = ",;"

	tokenizeString := func(str, delimiters string) []string {
		return strings.FieldsFunc(str, func(r rune) bool {
			return strings.ContainsRune(delimiters, r)
		})
	}

	tokens := tokenizeString(line, delimiters)

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	for _, token := range tokens {
		for key := range colors {
			if strings.Contains(token, key) {
				colors[key] = max(colors[key], getNum(token))			
			}
		}
	}
	
	var cubePower int = 1
	for _, val := range colors {
		cubePower *= val
	}

	return cubePower
}

func solver_part1() {
	fileScanner, readFile := read()

	var (
		counterId int = 1
		sumId     int = 0
	)
	for fileScanner.Scan() {
		if isPossible_part1(fileScanner.Text()) {
			sumId += counterId
		}

		counterId++
	}

	fmt.Println(sumId)

	readFile.Close()
}

func solver_part2() {
	fileScanner, readFile := read()

	var sumPower int = 0
	for fileScanner.Scan() {
		sumPower += cubePowerPart2(fileScanner.Text())
	}

	fmt.Println(sumPower)

	readFile.Close()
}

func main() {
	solver_part2()
}