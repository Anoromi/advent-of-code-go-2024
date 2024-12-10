package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func advent8_1() {
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
				potential1 := a.add(a.subtract(b))
				potential2 := b.add(b.subtract(a))
				if testForAntinode(board, antinodePositions, potential1) {
					antinodePositions[potential1.y*len(board[0])+potential1.x] = true
          fmt.Printf("potential1: %v\n", potential1)
				}
				if testForAntinode(board, antinodePositions, potential2) {
					antinodePositions[potential2.y*len(board[0])+potential2.x] = true
          fmt.Printf("potential2: %v\n", potential2)
				}

			}
		}
	}

	fmt.Printf("sum: %v\n", len(antinodePositions))

}

func (a Point) subtract(b Point) Point {
	return Point{a.x - b.x, a.y - b.y}
}
func (a Point) add(b Point) Point {
	return Point{b.x + a.x, b.y + a.y}
}

func testForAntinode(board [][]rune, previousAntinodes map[int]bool, position Point) bool {
	if position.y >= len(board) || position.y < 0 {
		return false
	}
	if position.x >= len(board[0]) || position.x < 0 {
		return false
	}

	flattenedPosition := position.y*len(board[0]) + position.x
	return !previousAntinodes[flattenedPosition]

}
