package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func advent11_1() {
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

  for i := 0; i < 25; i++ {
    slice = calculateStoneChanges(slice)
    fmt.Printf("slice %v: %v\n", i + 1, slice)
  }


	fmt.Printf("sum: %v\n", len(slice))
}

func calculateStoneChanges(previous []int64) []int64 {
	nextSlice := make([]int64, 0, len(previous))

	for _, v := range previous {
		if v == 0 {
			nextSlice = append(nextSlice, 1)
		} else if length := numberLength(v); length%2 == 0 {
			leftPart := v
			rightPart := v
      accum := int64(1)
      for i := 0; i < length/2; i++ {
				leftPart /= 10
				accum *= 10
			}
      rightPart = v % accum
			nextSlice = append(nextSlice, leftPart, rightPart)
		} else {
			nextSlice = append(nextSlice, v*2024)
		}
	}
	return nextSlice
}

func numberLength(value int64) int {
	if value == 0 {
		return 1
	}
	size := 0
	for value > 0 {
		value /= 10
		size++
	}
	return size

}
