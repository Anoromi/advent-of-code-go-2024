package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
)

func advent6_2() {
	filePath := "input6.txt"
	// filePath := "input6_test.txt"
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	position := Point{-1, -1}
	direction := Point{-1, -1}

	guardMap := make([][]rune, 0)

	for i := 0; scanner.Scan(); i++ {
		text := []rune(scanner.Text())

		for j, v := range text {
			isPosition, potentialDirection := interpretDirection(v)
			if isPosition {
				position = Point{j, i}
				direction = potentialDirection
			}
		}
		guardMap = append(guardMap, text)

	}

	traversed := findTraversedPositions(position, direction, guardMap)

	sum := 0
	for v := range traversed {
		x := v % len(guardMap)
		y := (v - position.x) / len(guardMap)
		temp := guardMap[y][x]
		guardMap[y][x] = '#'
		if temp != '#' && hasALoop(position, direction, guardMap) {
			sum += 1
		}
		guardMap[y][x] = temp
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("sum: %v\n", sum)

}

func hasALoop(position Point, direction Point, guardMap [][]rune) bool {
	hitPost := make(map[int][]Point)

	for {
		moved := Point{position.x + direction.x, position.y + direction.y}
		if moved.y >= len(guardMap) || moved.y < 0 {
			break
		}
		if moved.x >= len(guardMap[0]) || moved.x < 0 {
			break
		}
		flattenedPosition := len(guardMap)*position.y + position.x

		if guardMap[moved.y][moved.x] == '#' && slices.Contains(hitPost[flattenedPosition], direction) {
			return true
		}
		if guardMap[moved.y][moved.x] == '#' {
			if hitPost[flattenedPosition] == nil {
				hitPost[flattenedPosition] = make([]Point, 0)
			}
			hitPost[flattenedPosition] = append(hitPost[flattenedPosition], direction)
			direction = turnRight(direction)
			continue
		}

		position = moved
	}
	return false
}
