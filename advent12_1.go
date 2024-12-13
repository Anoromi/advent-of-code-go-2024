package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func advent12_1() {
	filePath := "input12.txt"
	// filePath := "input12_test.txt"
	// filePath := "input12_test2.txt"
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	passageMap := make([][]rune, 0)

	usedPositions := make([][]bool, 0)

	for i := 0; scanner.Scan(); i++ {
		text := []rune(scanner.Text())
		passageMap = append(passageMap, text)
		usedPositions = append(usedPositions, make([]bool, len(text)))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	sum := 0

	for i := 0; i < len(passageMap); i++ {
		for j := 0; j < len(passageMap[i]); j++ {
			if !usedPositions[i][j] {
				fmt.Printf("i: %v\n", i)
				sum += expandRegion(passageMap, usedPositions, i, j)
			}
		}
	}

	fmt.Printf("sum: %v\n", sum)
}

func expandRegion(passageMap [][]rune, usedPositions [][]bool, i int, j int) int {
	searchedChar := passageMap[i][j]

	unresolvedValues := []Point{{j, i}}
	area := 0
	perimeter := 0

	for len(unresolvedValues) > 0 {
		next := unresolvedValues[0]
		unresolvedValues = unresolvedValues[1:]
		if usedPositions[next.y][next.x] {
			continue
		}

		usedPositions[next.y][next.x] = true

		perimiterAdded := 4

		if next.x > 0 && passageMap[next.y][next.x-1] == searchedChar {
			unresolvedValues = append(unresolvedValues, Point{next.x - 1, next.y})
			perimiterAdded--
		}

		if next.y > 0 && passageMap[next.y-1][next.x] == searchedChar {
			unresolvedValues = append(unresolvedValues, Point{next.x, next.y - 1})
			perimiterAdded--
		}

		if next.x < len(passageMap[0])-1 && passageMap[next.y][next.x+1] == searchedChar {
			unresolvedValues = append(unresolvedValues, Point{next.x + 1, next.y})
			perimiterAdded--
		}

		if next.y < len(passageMap)-1 && passageMap[next.y+1][next.x] == searchedChar {
			unresolvedValues = append(unresolvedValues, Point{next.x, next.y + 1})
			perimiterAdded--
		}
		area++
		perimeter += perimiterAdded
	}

	return area * perimeter

}
