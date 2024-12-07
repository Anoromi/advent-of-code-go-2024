package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func advent5_2() {
	filePath := "input5.txt"
	// filePath := "input5_test.txt"
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer f.Close()

	findUnorderedWithSorting(f)

}

func findUnorderedWithSorting(f *os.File) {
	scanner := bufio.NewScanner(f)
	comeAfter := make(map[int][]int, 0)

	for scanner.Scan() {
		next := scanner.Text()
		if next == "" {
			break
		}

		arr := strings.Split(next, "|")

		left, errL := strconv.ParseInt(arr[0], 10, 32)
		if errL != nil {
			log.Fatal(errL)
		}
		right, errR := strconv.ParseInt(arr[1], 10, 32)
		if errR != nil {
			log.Fatal(errL)
		}

		leftInt := int(left)
		rightInt := int(right)

		if comeAfter[rightInt] == nil {
			comeAfter[rightInt] = make([]int, 0)
		}
		comeAfter[rightInt] = append(comeAfter[rightInt], leftInt)

	}

	sum := 0
	for scanner.Scan() {
		next := scanner.Text()

		instruction, err := stringSliceToIntSlice(strings.Split(next, ","))
		if err != nil {
			log.Fatal(err)
		}

		correct := isSorted(instruction, comeAfter)
		if correct {
			continue
		}

		for !isSorted(instruction, comeAfter) {
		processing:
			for i := 0; i < len(instruction); i++ {

				for j := i + 1; j < len(instruction); j++ {
					previousInstruction := instruction[j]
					if slices.Contains(comeAfter[instruction[i]], previousInstruction) {
						temp := instruction[i]
						instruction[i] = instruction[j]
						instruction[j] = temp

						break processing
					}
				}
			}
		}
		sum += instruction[len(instruction)/2]

	}
	fmt.Printf("sum: %v\n", sum)
}

func isSorted(instruction []int, comeAfter map[int][]int) bool {

	correct := true
processing:
	for i := 0; i < len(instruction); i++ {

		for _, previousInstruction := range instruction[i+1:] {
			if slices.Contains(comeAfter[instruction[i]], previousInstruction) {
				correct = false
				break processing
			}
		}
	}
	return correct
}
