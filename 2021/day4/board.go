package main

import (
	"fmt"
	"strconv"
	"strings"
)

type board []int

func parseBoard(data []string) (board, error) {
	if len(data) != 5 {
		return nil, fmt.Errorf("invalid number of lines: %d", len(data))
	}

	var numbers board
	for _, line := range data {
		fields := strings.Fields(line)
		if len(fields) != 5 {
			return nil, fmt.Errorf("want five numbers, got %d", len(fields))
		}
		for _, f := range fields {
			n, err := strconv.Atoi(f)
			if err != nil {
				return nil, fmt.Errorf("reading numbers: %w", err)
			}
			numbers = append(numbers, n)
		}
	}

	return numbers, nil
}

func find(b board, number int) int {
	for n, num := range b {
		if num == number {
			return n
		}
	}
	return -1
}

func winAt(b board, drawn []int) (int, int) {
	var (
		rows = make(map[int]int)
		cols = make(map[int]int)
	)

	for i, num := range drawn {
		index := find(b, num)
		if index < 0 {
			continue
		}

		b[index] = -1
		col, row := index%5, index/5
		rows[row]++
		cols[col]++
		if rows[row] == 5 || cols[col] == 5 {
			val := 0
			for _, n := range b {
				if n >= 0 {
					val += n
				}
			}
			return val, i
		}
	}

	return -1, -1
}
