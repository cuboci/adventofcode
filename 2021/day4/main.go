package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	flag.Parse()
	if flag.NArg() < 1 {
		log.Fatal("missing input")
	}

	numbers, boards, err := read(flag.Arg(0))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("value of first board to win: %d\n", firstWinResult(boards, numbers))
	fmt.Printf("value of last board to win: %d\n", lastWinResult(boards, numbers))
}

func read(file string) ([]int, []board, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, nil, fmt.Errorf("opening input file: %w", err)
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	if !sc.Scan() {
		return nil, nil, fmt.Errorf("empty input")
	}
	numbers, err := getNumbers(sc.Text())
	if err != nil {
		return nil, nil, err
	}

	var (
		boardData []string
		boards    []board
		b         board
	)
	for sc.Scan() {
		line := sc.Text()
		if strings.TrimSpace(line) == "" {
			if boardData != nil {
				b, err = parseBoard(boardData)
				if err != nil {
					return nil, nil, err
				}
				boards = append(boards, b)
			}
			boardData = nil
			continue
		}
		boardData = append(boardData, line)
	}
	boards = append(boards, b)
	if err := sc.Err(); err != nil {
		return nil, nil, fmt.Errorf("reading boards: %w", err)
	}

	return numbers, boards, nil
}

func getNumbers(input string) ([]int, error) {
	var numbers []int
	for _, val := range strings.Split(input, ",") {
		num, err := strconv.Atoi(val)
		if err != nil {
			return nil, fmt.Errorf("converting numbers: %w", err)
		}
		numbers = append(numbers, num)
	}

	return numbers, nil
}

func firstWinResult(boards []board, numbers []int) int {
	minNum := -1
	winValue := -1
	for _, b := range boards {
		btmp := make(board, len(b))
		copy(btmp, b)
		val, num := winAt(btmp, numbers)
		if num < 0 {
			continue
		}
		if minNum == -1 || num < minNum {
			minNum = num
			winValue = val
		}
	}

	return numbers[minNum] * winValue
}

func lastWinResult(boards []board, numbers []int) int {
	minNum := -1
	winValue := -1
	for _, b := range boards {
		btmp := make(board, len(b))
		copy(btmp, b)
		val, num := winAt(btmp, numbers)
		if num > minNum {
			minNum = num
			winValue = val
		}
	}

	return winValue * numbers[minNum]
}
