package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func advent4_2() {
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


	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[i]); j++ {
      aDiagonal := checkMas(lines, i - 1, j - 1, 1, 1) || checkMas(lines, i + 1, j + 1, -1, -1)
      bDiagonal := checkMas(lines, i + 1, j - 1, -1, 1) || checkMas(lines, i - 1, j + 1, 1, -1)
      if aDiagonal && bDiagonal {
        sum += 1
      }
		}
	}
		fmt.Printf("sum: %v\n", sum)
}


func checkMas(board [][]rune, startX int, startY int, directionX int, directionY int) bool {
	if startY >= len(board[0]) || startY < 0 {
		return false
	}
	if startY+directionY*2 >= len(board[0]) || startY + directionY * 2 < 0 {
		return false
	}
	if startX >= len(board) || startX < 0 {
		return false
	}
	if startX+directionX*2 >= len(board) || startX + directionX * 2 < 0 {
		return false
	}

	if board[startX][startY] != 'M' {
		return false
	}
	if board[startX+directionX*1][startY+directionY*1] != 'A' {
		return false
	}
	if board[startX+directionX*2][startY+directionY*2] != 'S' {
		return false
	}
	return true
}
