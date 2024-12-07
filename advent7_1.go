package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func advent7_1() {
	filePath := "input7.txt"
	// filePath := "input7_test.txt"
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var sum int64 = 0
	for i := 0; scanner.Scan(); i++ {
		text := scanner.Text()

		colonSplit := strings.Split(text, ":")

		expected, err := strconv.ParseInt(colonSplit[0], 10, 64)
		if err != nil {
			log.Fatal(err)
		}

		expectedSum := expected

		strParams := strings.Split(strings.Trim(colonSplit[1], " "), " ")

		params, err := stringSliceToIntSlice(strParams)
		if err != nil {
			log.Fatal(err)
		}

		if canCalculate(expectedSum, 0, params) {
			sum += expectedSum
		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("sum: %v\n", sum)

}

func canCalculate(value int64, current int64, left []int) bool {
	if current > value {
		return false
	}
	if len(left) == 0 {
		return current == value
	}

	next := int64(left[0])

	return canCalculate(value, current+next, left[1:]) ||
		canCalculate(value, current*next, left[1:]) ||
		canCalculate(value, concatenate(current, next), left[1:])
}

func concatenate(a, b int64) int64 {
	// var factor int64 = 0
  result := a
	{
		temp := b
		if temp == 0 {
      result *= 10
		} else {
			for temp > 0 {
				temp /= 10
        result *= 10
			}
		}
	}
  // fmt.Printf("a: %v\n", a)
  // fmt.Printf("b: %v\n", b)
  // fmt.Printf("factor: %v\n", factor)

	return result + b

}
