package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func advent11_2() {
	filePath := "input11.txt"
	// filePath := "input11_test.txt"
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	scanned := scanner.Scan()

	if err := scanner.Err(); err != nil || !scanned {
		log.Fatal(err)
	}

	text := strings.Split(scanner.Text(), " ")
	slice, err := stringSliceToInt64Slice(text)
	if err != nil {
		log.Fatal(err)
	}

	memo := make(map[SplitInfo]int64)

	sum := int64(0)
	for _, v := range slice {
		fmt.Printf("v: %v\n", v)
		sum += calculateSplittingCountForNumber(memo, v, 75)
	}

	fmt.Printf("sum: %v\n", sum)
}

type SplitInfo struct {
	value int64
	depth int
}

func calculateSplittingCountForNumber(memo map[SplitInfo]int64, value int64, depth int) int64 {
	if depth == 0 {
		return 1
	} else if v, exists := memo[SplitInfo{value, depth}]; exists {
		return v
	}

	if value == 0 {
		result := calculateSplittingCountForNumber(memo, 1, depth-1)
		memo[SplitInfo{1, depth - 1}] = result
		return result
	} else if length := numberLength(value); length%2 == 0 {
		leftPart := value
		rightPart := value
		accum := int64(1)
		for i := 0; i < length/2; i++ {
			leftPart /= 10
			accum *= 10
		}
		rightPart = value % accum

		leftResult := calculateSplittingCountForNumber(memo, leftPart, depth-1)
		memo[SplitInfo{leftPart, depth - 1}] = leftResult

		rightResult := calculateSplittingCountForNumber(memo, rightPart, depth-1)
		memo[SplitInfo{rightPart, depth - 1}] = rightResult
		return leftResult + rightResult
	} else {
		result := calculateSplittingCountForNumber(memo, value*2024, depth-1)
		memo[SplitInfo{value * 2024, depth - 1}] = result
		return result
	}

}
