package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func read() (*bufio.Scanner, *os.File) {
	readFile, err := os.Open("Day06/Day06_part1/6_1.txt")

	if err != nil {
		fmt.Print(err)
	}

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	return fileScanner, readFile
}

func getNums(line string) int {
	re := regexp.MustCompile(`\d+`)
	numbers := re.FindAllString(line, -1)

	r, _ := strconv.Atoi(numbers[0] + numbers[1] + numbers[2] + numbers[3])

	return r
}

type pair struct {
	time int
	distance int
}

func compute(a, b int) int {
	count := 0
	n := a
	for i := 1; i <= a; i++ {
		if i * (n - i) > b {
			count++
		}
	}
	return count
}

func ans(a pair) int {
	return compute(a.time, a.distance)
}

func solver() {
	fileScanner, readFile := read()
	defer readFile.Close()

	fileScanner.Scan()
	time := getNums(fileScanner.Text())
	
	fileScanner.Scan()
	distance := getNums(fileScanner.Text())

	fmt.Println(ans(pair{time, distance}))
}

func main() {
	solver()
}
