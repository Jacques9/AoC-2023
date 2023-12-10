package main

import (
	"bufio"
	"fmt"
	"os"
)

func read() (*bufio.Scanner, *os.File) {
	readFile, err := os.Open("Day10/Day10_part1/10_1.txt")

	if err != nil {
		fmt.Print(err)
	}

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	return fileScanner, readFile
}

func shortestPath(maze [][]rune) int {
	type Point struct {
		x, y int
	}

	var (
		nextPoints = func(maze [][]rune, p Point) []Point {
			points := []Point{}

			switch maze[p.y][p.x] {
			case '|':
				points = append(points, Point{p.x, p.y + 1})
				points = append(points, Point{p.x, p.y - 1})
			case '-':
				points = append(points, Point{p.x + 1, p.y})
				points = append(points, Point{p.x - 1, p.y})
			case 'L':
				points = append(points, Point{p.x, p.y - 1})
				points = append(points, Point{p.x + 1, p.y})
			case 'J':
				points = append(points, Point{p.x, p.y - 1})
				points = append(points, Point{p.x - 1, p.y})
			case '7':
				points = append(points, Point{p.x, p.y + 1})
				points = append(points, Point{p.x - 1, p.y})
			case 'F':
				points = append(points, Point{p.x, p.y + 1})
				points = append(points, Point{p.x + 1, p.y})
			case '.':
			case 'S':
				up := maze[p.y - 1][p.x]
				down := maze[p.y + 1][p.x]
				right := maze[p.y][p.x + 1]
				left := maze[p.y][p.x - 1]

				if down == '|' ||
					down == 'L' ||
					down == 'J' {
					points = append(points, Point{p.x, p.y + 1})
				}

				if right == '-' ||
					right == '7' ||
					right == 'J' {
					points = append(points, Point{p.x + 1, p.y})
				}

				if up == '|' ||
					up == '7' ||
					up == 'F' {
					points = append(points, Point{p.x, p.y - 1})
				}

				if left == '-' ||
					left == 'L' ||
					left == 'F' {
					points = append(points, Point{p.x - 1, p.y})
				}
			}

			return points
		}

		visited     = make(map[Point]int)
		queue       = []Point{}
		maxDistance = 0
		rows, cols  = len(maze), len(maze[0])
	)

	var startPoint Point
	for x := 0; x < rows; x++ {
		for y := 0; y < cols; y++ {
			if maze[x][y] == 'S' {
				startPoint = Point{x, y}
				queue = append(queue, startPoint)
				break
			}
		}
	}

	visited[startPoint] = 0

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		next := nextPoints(maze, current)

		for _, point := range next {
			if _, found := visited[point]; !found {
				visited[point] = visited[current] + 1
				maxDistance = max(maxDistance, visited[current] + 1)
				queue = append(queue, point)
			}
		}
	}
	return maxDistance
}

func solver() {
	fileScanner, readFile := read()

	defer readFile.Close()

	var maze [][]rune
	for fileScanner.Scan() {
		line := fileScanner.Text()

		arr := make([]rune, 0)
		for _, char := range line {
			arr = append(arr, char)
		}

		maze = append(maze, arr)
	}

	fmt.Println(shortestPath(maze))
}

func main() {
	solver()
}
