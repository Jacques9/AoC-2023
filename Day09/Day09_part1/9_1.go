package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

func read() (*bufio.Scanner, *os.File) {
	readFile, err := os.Open("Day09/Day09_part1/9_1.txt")

	if err != nil {
		fmt.Print(err)
	}

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	return fileScanner, readFile
}

func getArr(line string) []int {
	strs := strings.Split(line, " ")
	arr := make([]int, len(strs))

	for i := range arr {
		arr[i], _ = strconv.Atoi(strs[i])
	}
	return arr
}

func getNum(arr []int) int {
	var (
		value int = arr[len(arr) - 1]
		hasAllZero = func (arr []int) bool {
			for _, val := range arr {
				if val != 0 {
					return false
				}
			}
		
			return true
		}
	)
	
	for !hasAllZero(arr) {
		var v []int
		if len(arr) == 1 {
			value += arr[0]
			break
		}	

		for i := 0; i < len(arr) - 1; i++ {
			v = append(v, arr[i + 1] - arr[i])
			
		}
		fmt.Println(v)
		value += v[len(v) - 1]
		arr = v
	}
	return value
}

func solver() {
	fileScanner, readFile := read()
	
	defer readFile.Close()

	var sumVal int = 0
	for fileScanner.Scan() {
		sumVal += getNum(getArr(fileScanner.Text()))
	}

	fmt.Println(sumVal)
}

func main() {
	solver()
}