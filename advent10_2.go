package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func advent10_2() {
	filePath := "input10.txt"
	// filePath := "input10_test.txt"
	// filePath := "input10_test2.txt"
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	topMap := make([][]int, 0)

	trailheads := make([]Point, 0)

	for i := 0; scanner.Scan(); i++ {
		text := []rune(scanner.Text())
		values := make([]int, 0, len(text))
		for _, v := range text {
			v := int(v - '0')
			if v == '.' {
				v = -1
			}
			values = append(values, v)
		}
		topMap = append(topMap, values)
		for j := 0; j < len(text); j++ {
			if text[j] == '0' {
				trailheads = append(trailheads, Point{j, i})
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	sum := 0
	for _, v := range trailheads {
		sum += calculateTrailheadRatings(topMap, v)
	}

	fmt.Printf("sum: %v\n", sum)
}

func calculateTrailheadRatings(board [][]int, start Point) int {
	points := []Point{start}

	k := 0
	for len(points) > 0 {
		next := points[0]
		expectedValue := board[next.y][next.x] + 1
		points = points[1:]

		if board[next.y][next.x] == 9 {
			k++
			continue
		}

		if next.y > 0 && board[next.y-1][next.x] == expectedValue {
			points = append(points, Point{next.x, next.y - 1})
		}
		if next.x > 0 && board[next.y][next.x-1] == expectedValue {
			points = append(points, Point{next.x - 1, next.y})
		}
		if next.y < len(board)-1 && board[next.y+1][next.x] == expectedValue {
			points = append(points, Point{next.x, next.y + 1})
		}
		if next.x < len(board[0])-1 && board[next.y][next.x+1] == expectedValue {
			points = append(points, Point{next.x + 1, next.y})
		}
	}

	return k
}
