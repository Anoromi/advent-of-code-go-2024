package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
)

func advent9_2() {
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

	partitions := make([]PartitionInfo, 0)

	for i := 0; i < len(text); i += 1 {
		if text[i] < '0' || text[i] > '9' {
			log.Fatal(text[i])
		}
		size := int(text[i] - '0')
		var partition PartitionInfo
		if i%2 == 0 {
			block := []DataBlock{{i / 2, size}}
			partition = PartitionInfo{block, 0}
		} else {
			partition = PartitionInfo{nil, size}
		}
		partitions = append(partitions, partition)
	}

	sum := 0
	for {
		foundMove := false
		jStart := len(partitions) - 1
		if jStart%2 != 0 {
			jStart--
		}
		for j := jStart; j >= 0; j -= 2 {
			if len(partitions[j].filled) == 0 {
				continue
			}

			jPartition := &partitions[j]
			for i := 0; i < j; i++ {
				iPartition := &partitions[i]
				if iPartition.free >= jPartition.filled[0].size {
					iPartition.free -= jPartition.filled[0].size
					jPartition.free += jPartition.filled[0].size
					iPartition.filled = append(iPartition.filled, jPartition.filled[0])
					jPartition.filled = slices.Delete(jPartition.filled, 0, 1)
					foundMove = true
					break
				}
			}
		}
		if !foundMove {
			break
		}
	}

	accumulator := 0
	for _, v := range partitions {
		for _, k := range v.filled {
			for i := 0; i < k.size; i++ {
				sum += k.id * accumulator
				accumulator++
			}
		}
		accumulator += v.free
	}

	fmt.Printf("sum: %v\n", sum)

}

type PartitionInfo struct {
	filled []DataBlock
	free   int
}

type DataBlock struct {
	id   int
	size int
}
