package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func advent4_1() {
	filePath := "input4.txt"
	// filePath := "input4_test.txt"
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	lines := make([][]rune, 0)

	for scanner.Scan() {
		next := scanner.Text()

		characters := make([]rune, 0)
		for _, v := range next {
			characters = append(characters, v)

		}

		lines = append(lines, characters)
	}

	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}

	sum := 0
  fmt.Printf("lines: %v\n", lines)

	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[i]); j++ {
			count := boolToInt(checkXMAS(lines, i, j, 0, 1)) +
				boolToInt(checkXMAS(lines, i, j, 0, -1)) +
				boolToInt(checkXMAS(lines, i, j, 1, 0)) +
				boolToInt(checkXMAS(lines, i, j, -1, 0)) +
				boolToInt(checkXMAS(lines, i, j, 1, 1)) +
				boolToInt(checkXMAS(lines, i, j, -1, 1)) +
				boolToInt(checkXMAS(lines, i, j, -1, -1)) +
				boolToInt(checkXMAS(lines, i, j, 1, -1))

			sum += count
		}
	}
		fmt.Printf("sum: %v\n", sum)
}
func boolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}

func checkXMAS(board [][]rune, startX int, startY int, directionX int, directionY int) bool {
	if startY >= len(board[0]) || startY < 0 {
		return false
	}
	if startY+directionY*3 >= len(board[0]) || startY + directionY * 3 < 0 {
		return false
	}
	if startX >= len(board) || startX < 0 {
		return false
	}
	if startX+directionX*3 >= len(board) || startX + directionY * 3 < 0 {
		return false
	}

	if board[startX][startY] != 'X' {
		return false
	}
	if board[startX+directionX*1][startY+directionY*1] != 'M' {
		return false
	}
	if board[startX+directionX*2][startY+directionY*2] != 'A' {
		return false
	}
	if board[startX+directionX*3][startY+directionY*3] != 'S' {
		return false
	}
	return true
}
