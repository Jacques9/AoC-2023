package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

func read() (*bufio.Scanner, *os.File) {
	readFile, err := os.Open("Day09/Day09_part2/9_2.txt")

	if err != nil {
		fmt.Print(err)
	}

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	return fileScanner, readFile
}

func getArr(line string) []int {

	var (
		strs = strings.Split(line, " ")
		arr = make([]int, len(strs))
		reverse = func(arr []int) {
			n := len(arr)
			for i := 0; i < n / 2; i++ {
				arr[i], arr[n - i - 1] = arr[n - i - 1], arr[i]
			}
		}
	)

	for i := range arr {
		arr[i], _ = strconv.Atoi(strs[i])
	}

	reverse(arr)
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
	
	var res []int

	for !hasAllZero(arr) {
		var v []int
		if len(arr) == 1 {
			value -= arr[0]
			res = append(res, arr[0])
			break
		}
		res = append(res, arr[len(arr) - 1])
		for i := 0; i < len(arr) - 1; i++ {
			v = append(v, arr[i] - arr[i + 1])
			
		}
		value -= v[len(v) - 1]

		arr = v
	}
	
	val := res[len(res) - 1]
	for i := len(res) - 1; i > 0; i-- {
		val = res[i - 1] - val
	}
	
	return val
}

func solver() {
	fileScanner, readFile := read()
	
	defer readFile.Close()

	var sumVal int = 0
	for fileScanner.Scan() {
		sumVal += getNum((getArr(fileScanner.Text())))
	}

	fmt.Println(sumVal)
}

func main() {
	solver()
}