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

func getNums(line string) (a, b, c, d int) {
	re := regexp.MustCompile(`\d+`)
	numbers := re.FindAllString(line, -1)

	a, _ = strconv.Atoi(numbers[0])
	b, _ = strconv.Atoi(numbers[1])
	c, _ = strconv.Atoi(numbers[2])
	d, _ = strconv.Atoi(numbers[3])

	return a, b, c, d
}

type pair struct {
	time int
	distance int
}

func compute(a, b int) int {
	count := 0
	n := a
	for i := 1; i <= a; i++ {
		if (n - (n - i)) * (n - i) > b {
			count++
		}
	}
	return count
}

func ans(a, b, c, d pair) int {
	return compute(a.time, a.distance) *
		   compute(b.time, b.distance) *
		   compute(c.time, c.distance) *
		   compute(d.time, d.distance)
}

func solver() {
	fileScanner, readFile := read()
	defer readFile.Close()

	fileScanner.Scan()
	time_a, time_b, time_c, time_d := getNums(fileScanner.Text())
	
	fileScanner.Scan()
	distance_a, distance_b, distance_c, distance_d := getNums(fileScanner.Text())

	fmt.Println(ans(pair{time_a, distance_a}, pair{time_b, distance_b}, pair{time_c, distance_c}, pair{time_d, distance_d}))
}

func main() {
	solver()
}
