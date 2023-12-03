package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	// "golang.org/x/tools/go/ssa/interp"
)

func read() (*bufio.Scanner, *os.File) {
	readFile, err := os.Open("Day03/3.txt")

	if err != nil {
		fmt.Print(err)
	}

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	return fileScanner, readFile
}

func getArr(line string) []string {
	arr := []string {}
	for _, char := range line {
		arr = append(arr, string(char))
	}

	return arr
}

func readArr(arr [][]string) {
	for _, v := range arr {
		for _, el := range v {
			fmt.Print(el)
		}
		fmt.Println()
	}
}

func fetchDataObjPerLine(line []string) map[int]map[string]interface{} {
	var (
		data = map[int]map[string]interface{}{}
		getNum = func(str string) int {
			var num int = 0
			for _, char := range str {
				digit := int(char - '0')
				num = num * 10 + digit
			}

			return num
		}	
		isNum = func(char string) bool {
			return "0" <= char && char <= "9"
		}
	)

	i := 0
	for i < len(line) {
		if isNum(line[i]) {
			j := i
			for j < len(line) && isNum(line[j]) {
				j++
			}
			num := getNum(strings.Join(line[i:j], ""))
			// fmt.Println(i, j, num)
			data[len(data) + 1] = map[string]interface{} {
				"number": num,
				"Range": Range{start: i, end: j - 1},
			}

			i = j
		}
		i++
	}

	return data
}

type Range struct {
	start, end int
}

func check(arr [][]string, indexLine, start, end int) bool {
	const toBeCompared string = "0123456789."

	if start > 0 && !strings.Contains(toBeCompared, arr[indexLine][start - 1]) || end < len(arr) - 1 && !strings.Contains(toBeCompared, arr[indexLine][end + 1]) {
		return true
	}	

	var (
		diagonalDownLeft bool = true
		diagonalDownRight bool = true
		diagonalUpLeft bool = true
		diagonalUpRight bool = true
	)

	if indexLine == 0 {
		diagonalUpLeft = false
		diagonalUpRight = false
	}

	if indexLine == len(arr) - 1 {
		diagonalDownLeft = false
		diagonalDownRight = false
	}

	if start == 0 {
		diagonalUpLeft = false
		diagonalDownLeft = false
	}

	if end == len(arr[indexLine]) - 1 {
		diagonalDownRight = false
		diagonalUpRight = false
	}

	for i := start; i <= end; i++ {
		if indexLine != 0 && !strings.Contains(toBeCompared, arr[indexLine - 1][i]) {
			return true
		}

		if indexLine != len(arr) - 1 && !strings.Contains(toBeCompared, arr[indexLine + 1][i]) {
			return true
		}
	}

	diagonalDownLeft = diagonalDownLeft && !strings.Contains(toBeCompared, arr[indexLine + 1][start - 1])
	diagonalDownRight = diagonalDownRight && !strings.Contains(toBeCompared, arr[indexLine + 1][end + 1])
	diagonalUpRight = diagonalUpRight && !strings.Contains(toBeCompared, arr[indexLine - 1][end + 1])
	diagonalUpLeft = diagonalUpLeft && !strings.Contains(toBeCompared, arr[indexLine - 1][start - 1])

	return diagonalDownLeft || diagonalDownRight || diagonalUpRight || diagonalUpLeft
}

func parseLine(dataObj map[int]map[string]interface{}, arr [][]string, indexLine int) int {
	var sumPerLine int = 0
	for _, innerMap := range dataObj {
		var (
			number = innerMap["number"]
			start =  innerMap["Range"].(Range).start
			end = innerMap["Range"].(Range).end
		)
		if check(arr, indexLine, start, end) {
			sumPerLine += number.(int)
		}
	}
	return sumPerLine
}

func parse(arr [][]string) int {
	var sumOverAll int = 0
	for index, line := range arr {
		sumOverAll += parseLine(fetchDataObjPerLine(line), arr, index)
	}

	return sumOverAll
}

func solver_part1() {
	fileScanner, readFile := read()

	arr := [][]string {}
	for fileScanner.Scan() {
		// fmt.Println(fileScanner.Text())
		arr = append(arr, getArr(fileScanner.Text()))
	}

	fmt.Println(parse(arr))

	readFile.Close()
}

func main() {
	solver_part1()
}