package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func read() (*bufio.Scanner, *os.File) {
	readFile, err := os.Open("Day01/1.txt")

	if err != nil {
		fmt.Print(err)
	}

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	return fileScanner, readFile
}

func getStringWithNums(str string) string {
	digits := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}

    var output string
	i := 0
	for i < len(str) {
		found := false

		for j := i + 1; j <= len(str); j++ {
			substr := str[i:j]
			if val, exists := digits[substr]; exists {
				output += val
				i = j - 1
				found = true
				break
			}
		}

		if !found {
			output += string(str[i])
			i++
		}
	}

    fmt.Println(output)

	return output
}


func fetchNum(str string) int {
    r := regexp.MustCompile("[^0-9]+").ReplaceAllString(getStringWithNums(str), "")
    // fmt.Println(r)

	if len(str) == 0 {
		return 0
	}

	num, err := strconv.Atoi(string(r[0]) + string(r[len(r) - 1]))

	if err != nil {
		fmt.Println("Error: ", err)
	}

	return num
}

func solver() {
	fileScanner, readFile := read()

	var sum int = 0
	for fileScanner.Scan() {
		// fmt.Println(fileScanner.Text())
		sum += fetchNum(fileScanner.Text())
	}

	fmt.Println(sum)

	readFile.Close()
}

func main() {
	solver()
}
