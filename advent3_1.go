package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func advent3_2() {
	filePath := "input3.txt"
	// filePath := "input3_test.txt"
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer f.Close()

	reader := bufio.NewReader(f)

	sum := 0
	do := true
	for {
		found, result, err := advance(reader)

		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		if found == Mul && do {
			sum += result
		} else if found == Dodont {
			switch result {
			case 0:
				do = false
			case 1:
				do = true
			}
		}

	}

	fmt.Printf("sum: %v\n", sum)
}

const (
	Error  int = -1
	Mul        = 0
	Dodont     = 1
)

func advance(reader *bufio.Reader) (int, int, error) {
	char, _, err := reader.ReadRune()
	if char == 'm' {
		success, value, err := mul(reader)
		if !success || err != nil {
			return 0, 0, err
		}
		return Mul, value, nil
	} else if char == 'd' {
		valid, do, err := dodont(reader)
		if !valid || err != nil {
			return 0, 0, err
		}
		returning := 0
		if do {
			returning = 1
		}
		return Dodont, returning, nil

	}

	return -1, 0, err
}

func dodont(reader *bufio.Reader) (valid bool, do bool, err error) {
	char, _, err := reader.ReadRune()
	if err != nil || char != 'o' {
		return false, false, err
	}
	char, _, err = reader.ReadRune()
	if err != nil {
		return false, false, err
	}
	if char == 'n' {
		char, _, err := reader.ReadRune()
		if err != nil || char != '\'' {
			return false, false, err
		}
		char, _, err = reader.ReadRune()
		if err != nil || char != 't' {
			return false, false, err
		}
		char, _, err = reader.ReadRune()
		if err != nil || char != '(' {
			return false, false, err
		}
		char, _, err = reader.ReadRune()
		if err != nil || char != ')' {
			return false, false, err
		}
		return true, false, nil
	}
	if char == '(' {
		char, _, err := reader.ReadRune()
		if err != nil || char != ')' {
			return false, false, err
		}
		return true, true, nil
	}
	return false, false, nil
}

func mul(reader *bufio.Reader) (bool, int, error) {
	char, _, err := reader.ReadRune()
	if err != nil || char != 'u' {
		return false, 0, err
	}
	char, _, err = reader.ReadRune()
	if err != nil || char != 'l' {
		return false, 0, err
	}
	char, _, err = reader.ReadRune()
	if err != nil || char != '(' {
		return false, 0, err
	}
	valid, left_number, err := integer(reader)
	if !valid || err != nil {
		return valid, 0, err
	}
	char, _, err = reader.ReadRune()
	if err != nil || char != ',' {
		return false, 0, err
	}
	valid, right_number, err := integer(reader)
	if !valid || err != nil {
		return valid, 0, err
	}
	char, _, err = reader.ReadRune()
	if err != nil || char != ')' {
		return false, 0, err
	}
	return true, left_number * right_number, nil

}

func integer(reader *bufio.Reader) (valid bool, value int, err error) {
	readAnything := false
	number := 0
	nextC, _, err := reader.ReadRune()
	for ; err == nil && (nextC >= '0' && nextC <= '9'); nextC, _, err = reader.ReadRune() {
		number = number*10 + int(nextC-48)
		readAnything = true
	}

	if err != nil {
		return false, 0, err
	}

	if reader.UnreadRune() != nil {
		return false, 0, err
	}

	if !readAnything {
		return false, 0, nil
	}
	return true, number, err
}
