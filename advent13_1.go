package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func advent13_1() {
	filePath := "input13.txt"
	// filePath := "input13_test.txt"
	// filePath := "input13_test2.txt"
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	sum := int64(0)

	for i := 0; scanner.Scan(); i++ {
		buttonAText := scanner.Text()

		if !scanner.Scan() {
			break
		}
		buttonBText := scanner.Text()

		if !scanner.Scan() {
			break
		}
		prizeText := scanner.Text()

		buttonAStrings := strings.Split(strings.Trim(strings.Split(buttonAText, ":")[1], " "), ",")
		axString := string([]rune(buttonAStrings[0])[2:])
		ayString := string([]rune(buttonAStrings[1])[2:])

		buttonBStrings := strings.Split(strings.Trim(strings.Split(buttonBText, ":")[1], " "), ",")
		bxString := string([]rune(buttonBStrings[0])[2:])
		byString := string([]rune(buttonBStrings[1])[2:])

		priceStrings := strings.Split(strings.Trim(strings.Split(prizeText, ":")[1], " "), ",")
		pricexString := string([]rune(priceStrings[0])[2:])
		priceyString := string([]rune(priceStrings[1])[3:])

    scanner.Scan()

		ax, err := strconv.ParseInt(axString, 10, 32)
		if err != nil {
			log.Panic(err)
		}
		ay, err := strconv.ParseInt(ayString, 10, 32)
		if err != nil {
			log.Panic(err)
		}

		bx, err := strconv.ParseInt(bxString, 10, 32)
		if err != nil {
			log.Panic(err)
		}
		by, err := strconv.ParseInt(byString, 10, 32)
		if err != nil {
			log.Panic(err)
		}

		pricex, err := strconv.ParseInt(pricexString, 10, 32)
		if err != nil {
			log.Panic(err)
		}

		markUp := int64(10000000000000)
		pricex += markUp
		pricey, err := strconv.ParseInt(priceyString, 10, 32)
		if err != nil {
			log.Panic(err)
		}
		pricey += markUp

		right := (ax*pricey - ay*pricex) / (ax*by - ay*bx)
		left := (pricex - bx*right) / ax

		if ax*left+bx*right == pricex && ay*left+by*right == pricey {
			sum += left*3 + right
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("sum: %v\n", sum)
}

type Pair[V any] struct {
	a V
	b V
}
