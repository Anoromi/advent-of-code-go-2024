package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int))
}
func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type InverseIntHeap []int

func (h InverseIntHeap) Len() int           { return len(h) }
func (h InverseIntHeap) Less(i, j int) bool { return h[j] < h[i] }
func (h InverseIntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *InverseIntHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *InverseIntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

var _ heap.Interface = (*InverseIntHeap)(nil)

func advent1_1() {
	filePath := "input2.txt"
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	leftValues := &IntHeap{}
	heap.Init(leftValues)
	rightValues := &IntHeap{}
	heap.Init(rightValues)

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

		heap.Push(leftValues, int(left))
		heap.Push(rightValues, int(right))
	}

	sum := 0
	for leftValues.Len() > 0 {
		lowest := heap.Pop(leftValues).(int)
		highest := heap.Pop(rightValues).(int)

		fmt.Printf("Values %d %d\n", lowest, highest)
		if highest > lowest {
			sum += highest - lowest
		} else {
			sum += lowest - highest
		}
	}
	fmt.Printf("Sum %d\n", sum)
}
