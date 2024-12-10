package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func advent8_2() {
	filePath := "input8.txt"
	// filePath := "input8_test.txt"
	// filePath := "input8_test2.txt"
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	board := make([][]rune, 0)

	antennas := make(map[rune][]Point)

	for i := 0; scanner.Scan(); i++ {
		text := []rune(scanner.Text())

		board = append(board, text)
		for j := range text {
			if text[j] != '.' {
				antennas[text[j]] = append(antennas[text[j]], Point{j, i})
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	antinodePositions := make(map[int]bool)

	for _, v := range antennas {
		for ai, a := range v {
			for _, b := range v[ai+1:] {
				direction1 := a.subtract(b)
				direction2 := b.subtract(a)
				fillRay(board, antinodePositions, a, direction1)
				fillRay(board, antinodePositions, a, direction2)

			}
		}
	}

	fmt.Printf("sum: %v\n", len(antinodePositions))

}

func fillRay(board [][]rune, antinodePositions map[int]bool, start Point, direction Point) {
	position := start
	for {
		if position.y >= len(board) || position.y < 0 {
			break
		}
		if position.x >= len(board[0]) || position.x < 0 {
			break
		}
		antinodePositions[position.y*len(board[0])+position.x] = true
		position = position.add(direction)
	}

}

func (a Point) invert() Point {
	return Point{-a.x, -a.y}
}
