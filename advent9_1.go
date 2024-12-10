package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func advent9_1() {
	filePath := "input9.txt"
	// filePath := "input9_test.txt"
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

	text := []rune(scanner.Text())

	usedBlocks := make([]ContainerInfo, 0)

	for i := 0; i < len(text); i += 1 {
		if text[i] < '0' || text[i] > '9' {
			log.Fatal(text[i])
		}
		size := int(text[i] - '0')
		var block ContainerInfo
		if i%2 == 0 {
			block = ContainerInfo{size, 0}
		} else {
			block = ContainerInfo{0, size}
		}
		usedBlocks = append(usedBlocks, block)
	}

	sum := 0
	i, accumulator := 0, 0

	for j := len(usedBlocks) - 1; i < j; {
		iBlock := &usedBlocks[i]
		jBlock := &usedBlocks[j]

		if iBlock.filled > 0 {
			for ; iBlock.filled > 0; iBlock.filled-- {
				sum += (i / 2) * accumulator
				accumulator++
			}
			continue
		}
		if iBlock.free == 0 {
			i++
			continue
		}
		if jBlock.filled == 0 {
			j--
			continue
		}

		sum += (j / 2) * accumulator
		accumulator++

		jBlock.filled--
		iBlock.free--
	}
	{
		iBlock := &usedBlocks[i]
		for ; iBlock.filled > 0; iBlock.filled-- {
			sum += (i / 2) * accumulator
			accumulator++
		}

	}

	fmt.Printf("sum: %v\n", sum)

}

type ContainerInfo struct {
	filled int
	free   int
}
