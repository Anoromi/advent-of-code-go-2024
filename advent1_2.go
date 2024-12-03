package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func advent1_2() {
	filePath := "input2.txt"
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	leftValues := []int{}
  rightValues := make(map[int]int)
	//rightValues := &IntHeap{}

	for scanner.Scan() {
		space := regexp.MustCompile(` +`)
		arr := space.Split(scanner.Text(), -1)
		//arr := strings.Split(scanner.Text(), " ")

		left, err := strconv.ParseInt(arr[0], 10, 32)
		if err != nil {
			log.Fatal(err)
			return
		}
		right, err := strconv.ParseInt(arr[1], 10, 32)
		if err != nil {
			log.Fatal(err)
			return
		}

    leftValues = append(leftValues, int(left))
    rightValues[int(right)] += 1;
		//heap.Push(rightValues, int(right))
	}

	sum := 0
  for _, left := range leftValues {
    sum += left * rightValues[left]
	}
	fmt.Printf("Sum %d\n", sum)
}
