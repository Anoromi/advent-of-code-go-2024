package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func advent12_2() {
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

	usedPositions := make([][]Crossed, 0)

	for i := 0; scanner.Scan(); i++ {
		text := []rune(scanner.Text())
		passageMap = append(passageMap, text)
		usedPositions = append(usedPositions, make([]Crossed, len(text)))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	sum := 0

	for i := 0; i < len(passageMap); i++ {
		for j := 0; j < len(passageMap[i]); j++ {
			if !usedPositions[i][j].vertically || !usedPositions[i][j].horizontally {
				sum += expandRegionWithSides(passageMap, usedPositions, i, j)
			}
		}
	}

	fmt.Printf("sum: %v\n", sum)
}

type Crossed struct {
	vertically   bool
	horizontally bool
}

func expandRegionWithSides(passageMap [][]rune, usedPositions [][]Crossed, i int, j int) int {
	searchedChar := passageMap[i][j]

	unresolvedValues := []Point{{j, i}}
	area := 0
	perimeter := 0

	for len(unresolvedValues) > 0 {
		next := unresolvedValues[0]
		unresolvedValues = unresolvedValues[1:]
		checkChange := func(x int, y int) (inBounds bool, isWall bool) {
			if x < 0 || x >= len(passageMap[0]) {
				return false, false
			}
			if y < 0 || y >= len(passageMap) {
				return false, false
			}
			return true, passageMap[y][x] != searchedChar
		}

		if !usedPositions[next.y][next.x].horizontally {

			start := next.x

			for ; start >= 0; start-- {
				if passageMap[next.y][start] != searchedChar {
					break
				}
			}
			start++

			previousTopHasWall := false
			previousBottomHasWall := false
			wallCount := 0

			if next.y == 0 {
				wallCount++
			}

			if next.y == len(passageMap)-1 {
				wallCount++
			}

			for i := start; i < len(passageMap[0]) && passageMap[next.y][i] == searchedChar; i++ {
				_, topHasWall := checkChange(i, next.y-1)
				if topHasWall && !previousTopHasWall {
					wallCount++
				}
				previousTopHasWall = topHasWall

				_, bottomHasWall := checkChange(i, next.y+1)
				if bottomHasWall && !previousBottomHasWall {
					wallCount++
				}
				previousBottomHasWall = bottomHasWall
				if !usedPositions[next.y][i].vertically && !usedPositions[next.y][i].horizontally {
					area++
				}
				usedPositions[next.y][i].horizontally = true
				unresolvedValues = append(unresolvedValues, Point{i, next.y})
			}
			perimeter += wallCount
		}
		if !usedPositions[next.y][next.x].vertically {
			start := next.y
			for ; start >= 0; start-- {
				if passageMap[start][next.x] != searchedChar {
					break
				}
			}
			start++

			previousLeftWasWall := false
			previousRightHasWall := false
			wallCount := 0

			if next.x == 0 {
				wallCount++
			}

			if next.x == len(passageMap[0])-1 {
				wallCount++
			}

			for i := start; i < len(passageMap[0]) && passageMap[i][next.x] == searchedChar; i++ {
				_, leftHasWall := checkChange(next.x-1, i)
				if leftHasWall && !previousLeftWasWall {
					wallCount++
				}
				previousLeftWasWall = leftHasWall

				_, rightHasWall := checkChange(next.x+1, i)
				if rightHasWall && !previousRightHasWall {
					wallCount++
				}
				previousRightHasWall = rightHasWall
				unresolvedValues = append(unresolvedValues, Point{next.x, i})
				if !usedPositions[i][next.x].vertically && !usedPositions[i][next.x].horizontally {
					area++
				}
				usedPositions[i][next.x].vertically = true
			}
			perimeter += wallCount
		}
	}

	fmt.Printf("searchedChar: %v\n", string(searchedChar))
	fmt.Printf("perimeter: %v\n", perimeter)
	fmt.Printf("area: %v\n", area)

	return area * perimeter

}
