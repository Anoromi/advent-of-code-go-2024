package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func advent2_1() {
	filePath := "input2.txt"
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	count := 0
	for scanner.Scan() {

		level := strings.Split(scanner.Text(), " ")
		parsed_levels := make([]int, len(level))
		for i, v := range level {
			parsed_value, err := strconv.ParseInt(v, 10, 32)
			if err != nil {
				log.Fatal(err)
			}
			parsed_levels[i] = int(parsed_value)
		}

		ascending := parsed_levels[1] > parsed_levels[0]

		valid := true
		for i := 0; i < len(level)-1; i++ {
			switch {
			case parsed_levels[i] > parsed_levels[i+1] && ascending:
				valid = false
				break
			case ascending && (parsed_levels[i+1]-parsed_levels[i] < 1 || parsed_levels[i+1]-parsed_levels[i] > 3):
				valid = false
				break
			case !ascending && (parsed_levels[i]-parsed_levels[i+1] < 1 || parsed_levels[i]-parsed_levels[i+1] > 3):
				valid = false
				break
			}
		}
		if valid {
			count++
		}

	}

	fmt.Printf("count: %v\n", count)
}
