package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

func advent2_2() {
	filePath := "input2.txt"
	// filePath := "input2_test.txt"
	// filePath := "input2_test2.txt"
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

		valid := false
		for i := 0; i < len(parsed_levels); i++ {
			tmp := make([]int, len(parsed_levels))
			copy(tmp, parsed_levels)
			if evaluateCalm(slices.Delete(tmp, i, i+1)) {
				valid = true
				break
			}
		}
		if valid || evaluateCalm(parsed_levels) {
			count++
		}
		if (valid || evaluateCalm(parsed_levels)) != (evaluateFunky(parsed_levels)) {
			fmt.Printf("problematic level: %v\n", level)
			fmt.Printf("(valid || evaluateCalm(parsed_levels)): %v\n", (valid || evaluateCalm(parsed_levels)))
		}
		//k := evaluate(parsed_levels, -1)
		// fmt.Printf("k: %v\n", k)
		// if k == -1 {
		// 	count++
		// } else if evaluate(slices.Delete(parsed_levels, k, k+1), -1) == -1 || evaluate(slices.Delete(parsed_levels, 0, 0), -1) == -1 {
		// 	count++
		// }
		// if k == -1 {
		// 	count++
		// } else if evaluate(parsed_levels, k) == -1 {
		// 	count++
		// } else if evaluate(parsed_levels, 0) == -1 {
		// 	count++
		// }

		// for i := 0; i < len(level)-1; i++ {
		// 	switch {
		// 	case parsed_levels[i] > parsed_levels[i+1] && ascending:
		// 		valid = false
		// 		break
		// 	case ascending && (parsed_levels[i+1]-parsed_levels[i] < 1 || parsed_levels[i+1]-parsed_levels[i] > 3):
		// 		valid = false
		// 		break
		// 	case !ascending && (parsed_levels[i]-parsed_levels[i+1] < 1 || parsed_levels[i]-parsed_levels[i+1] > 3):
		// 		valid = false
		// 		break
		// 	}
		// }

		//no_issues := true
		//for i := 0; i < len(level)-1; i++ {
		//	good := compare(parsed_levels, i, i+1)

		//	if no_issues && !good {
		//		no_issues = false
		//		if i != 0 && i != len(parsed_levels)-2 {
		//			good_without := compare(parsed_levels, i, i+2)
		//			if !good_without {
		//				correct = false
		//				break
		//			} else if parsed_levels[i+2] > parsed_levels[i] {
		//				ascending_count++
		//			}

		//		}
		//	} else if !good {
		//		correct = false
		//		break
		//	} else if parsed_levels[i+1] > parsed_levels[i] {
		//		ascending_count++
		//	}
		//}

	}

	fmt.Printf("count: %v\n", count)
}

func evaluateCalm(levels []int) bool {
	for i := 0; i < len(levels)-1; i++ {
		next := i + 1
		if abs(levels[next]-levels[i]) < 1 || abs(levels[next]-levels[i]) > 3 {
			return false
		}
	}
	if !sort.SliceIsSorted(levels, func(i, j int) bool { return levels[i] < levels[j] }) && !sort.SliceIsSorted(levels, func(i, j int) bool { return levels[i] > levels[j] }) {
		return false
	}
	return true
}

func evaluateFunky(levels []int) bool {
	k := evaluate(levels, -1)
  // fmt.Printf("k: %v\n", k)
	// fmt.Printf("k: %v\n", k)
	// if k == -1 {
	// 	return true
	// } else if evaluate(slices.Delete(levels, k, k+1), -1) == -1 || evaluate(slices.Delete(levels, 0, 0), -1) == -1 {
	// 	return true
	// }
	if k == -1 {
		return true
	} else if evaluate(levels, k) == -1 || evaluate(levels, k - 1) == -1 {
		return true
	} else if evaluate(levels, 0) == -1 {
		return true
	}
	return false
}

func evaluate(levels []int, exclude int) int {
	ascending := levels[1] > levels[0]
	if exclude == 0 {
		ascending = levels[2] > levels[1]
	}

	if exclude == 1 {
		ascending = levels[2] > levels[0]
	}
	for i := 0; i < len(levels)-1; i++ {
		if i == exclude {
			continue
		}
		next := i + 1
		if next == exclude {
			next++
		}
		if next >= len(levels) {
			return -1
		}

		switch {
		case levels[i] > levels[next] && ascending:
			return next
		case levels[i] < levels[next] && !ascending:
			return next
		case abs(levels[next]-levels[i]) < 1 || abs(levels[next]-levels[i]) > 3:
			return next
		}
	}
	// sort.SliceStable(levels, func(i, j int) bool { return levels[i] < levels[j] })
	return -1
}

func abs(v int) int {
	if v < 0 {
		return -v
	}
	return v
}
