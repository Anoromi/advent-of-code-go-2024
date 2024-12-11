package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func advent5_1() {
	filePath := "input5.txt"
	// filePath := "input5_test.txt"
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer f.Close()

	findUnordered(f)

}

func findUnordered(f *os.File) {
	scanner := bufio.NewScanner(f)
	comeAfter := make(map[int][]int, 0)

	for scanner.Scan() {
		next := scanner.Text()
		if next == "" {
			break
		}

		arr := strings.Split(next, "|")

		fmt.Printf("arr: %v\n", arr)

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

		correct := true

		fmt.Printf("next: %v\n", next)
	processing:
		for i := 0; i < len(instruction); i++ {

			for _, previousInstruction := range instruction[i+1:] {
				if slices.Contains(comeAfter[instruction[i]], previousInstruction) {
					correct = false
					fmt.Printf("instruction[i]: %v\n", instruction[i])
					fmt.Printf("previousInstruction: %v\n", previousInstruction)
					break processing
				}
			}
		}
		if correct {
			sum += instruction[len(instruction)/2]
		}
	}
	fmt.Printf("sum: %v\n", sum)
}

func funkyFindUnordered(f *os.File) {
	scanner := bufio.NewScanner(f)

	comeAfter := make(map[int][]int, 0)

	for scanner.Scan() {
		next := scanner.Text()
		if next == "" {
			break
		}

		arr := strings.Split(next, "|")

		left, errL := strconv.ParseInt(arr[0], 10, 32)
		right, errR := strconv.ParseInt(arr[1], 10, 32)
		if err := errors.Join(errL, errR); err != nil {
			log.Fatal(err)
		}

		leftInt := int(left)
		rightInt := int(right)

		if comeAfter[rightInt] == nil {
			comeAfter[rightInt] = make([]int, 0)
		}
		comeAfter[rightInt] = append(comeAfter[rightInt], leftInt)

	}

	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}

	comeAfterExpanded := unwrapDependencies(comeAfter)
	fmt.Printf("comeAfterExpanded: %v\n", comeAfterExpanded)
	sum := 0
	for scanner.Scan() {
		next := scanner.Text()

		instruction, err := stringSliceToIntSlice(strings.Split(next, ","))
		if err != nil {
			log.Fatal(err)
		}

		correct := true

		fmt.Printf("next: %v\n", next)
	processing:
		for i := 0; i < len(instruction); i++ {

			for _, previousInstruction := range instruction[i+1:] {
				if comeAfterExpanded[instruction[i]][previousInstruction] {
					correct = false
					fmt.Printf("instruction[i]: %v\n", instruction[i])
					fmt.Printf("previousInstruction: %v\n", previousInstruction)
					break processing
				}
			}
		}
		if correct {
			sum += instruction[len(instruction)/2]
		}

	}
	fmt.Printf("sum: %v\n", sum)
}

func stringSliceToIntSlice(s []string) (result []int, err error) {
	r := make([]int, len(s))
	for i := 0; i < len(r); i++ {
		v, err := strconv.ParseInt(s[i], 10, 32)
		if err != nil {
			return nil, err
		}
		r[i] = int(v)
	}

	return r, nil
}

func stringSliceToInt64Slice(s []string) (result []int64, err error) {
	r := make([]int64, len(s))
	for i := 0; i < len(r); i++ {
		v, err := strconv.ParseInt(s[i], 10, 32)
		if err != nil {
			return nil, err
		}
		r[i] = v
	}

	return r, nil
}

func intesect(a []int, b []int) bool {
	k := make(map[int]bool)
	for _, v := range a {
		k[v] = true
	}

	for _, v := range b {
		if k[v] {
			return true
		}
	}
	return false
}

func unwrapDependencies(values map[int][]int) map[int]map[int]bool {
	allPoints := make(map[int]bool)
	for k, values := range values {
		allPoints[k] = true
		for _, value := range values {
			allPoints[value] = true
		}
	}
	s := make([]int, 0)
	before := make(map[int]map[int]bool)
	for k := range allPoints {
		v := values[k]
		if len(v) == 0 {
			s = append(s, k)
			before[k] = make(map[int]bool)
		}
	}

	for len(s) > 0 {
		fmt.Printf("s: %v\n", s)
		currentLowest := s[0]

		linked := make([]int, 0)

		for k, kAdjacency := range values {
			if slices.Contains(kAdjacency, currentLowest) {
				linked = append(linked, k)
				if before[k] == nil {
					before[k] = make(map[int]bool)
				}
				for v := range before[currentLowest] {
					before[k][v] = true
				}
				before[k][currentLowest] = true
			}
		}
		fmt.Printf("currentLowest: %v\n", currentLowest)
		fmt.Printf("linked: %v\n", linked)
		s = append(s, linked...)[1:]
	}

	return before
}

