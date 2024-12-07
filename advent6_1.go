package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Vector struct {
	x int
	y int
}

func advent6_1() {
	filePath := "input6.txt"
	// filePath := "input6_test.txt"
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	position := Vector{-1, -1}
	direction := Vector{-1, -1}

	guardMap := make([][]rune, 0)

	for i := 0; scanner.Scan(); i++ {
		text := []rune(scanner.Text())

		for j, v := range text {
			isPosition, potentialDirection := interpretDirection(v)
			if isPosition {
				position = Vector{j, i}
				direction = potentialDirection
			}
		}
		guardMap = append(guardMap, text)

	}


	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}


  traversed := findTraversedPositions(position, direction, guardMap)
	fmt.Printf("sum: %v\n", len(traversed))

}

func findTraversedPositions(position Vector, direction Vector, guardMap [][]rune) map[int]bool {
	traversed := make(map[int]bool)

	for {
		traversed[len(guardMap)*position.y+position.x] = true

		moved := Vector{position.x + direction.x, position.y + direction.y}
		// fmt.Printf("moved: %v\n", moved)
		// fmt.Printf("direction: %v\n", direction)
		if moved.y >= len(guardMap) || moved.y < 0 {
			break
		}
		if moved.x >= len(guardMap[0]) || moved.x < 0 {
			break
		}

		if guardMap[moved.y][moved.x] == '#' {
			// fmt.Printf("Turning right\n")
			direction = turnRight(direction)
			continue
		}

		position = moved
	}
  return traversed
}

func turnRight(v Vector) Vector {
	return Vector{v.y * -1, v.x}
}

func interpretDirection(c rune) (bool, Vector) {
	switch c {
	case '<':
		return true, Vector{-1, 0}
	case '>':
		return true, Vector{1, 0}
	case '^':
		return true, Vector{0, -1}
	case 'v':
		return true, Vector{0, 1}
	}
	return false, Vector{}
}
