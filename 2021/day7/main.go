package main

import (
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	flag.Parse()
	if flag.NArg() < 1 {
		log.Fatal("missing input")
	}

	crabs, err := read(flag.Arg(0))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("minimum fuel: %d\n", minFuel(crabs))
	fmt.Printf("actual minimum fuel: %d\n", realFuel(crabs))
}

func read(file string) ([]int, error) {
	data, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}

	tmp := strings.Split(strings.TrimSpace(string(data)), ",")
	crabs := make([]int, 0, len(tmp))
	for _, val := range tmp {
		d, err := strconv.Atoi(val)
		if err != nil {
			return nil, err
		}
		crabs = append(crabs, d)
	}

	return crabs, nil
}

func minFuel(crabs []int) int {
	sort.Ints(crabs)
	minFuel := -1
	for pos := crabs[0]; pos <= crabs[len(crabs)-1]; pos++ {
		var fuel int
		for _, p := range crabs {
			if p >= pos {
				fuel += (p - pos)
			} else {
				fuel += (pos - p)
			}
		}
		if minFuel < 0 || fuel < minFuel {
			minFuel = fuel
		}
	}

	return minFuel
}

func realFuel(crabs []int) int {
	sort.Ints(crabs)
	minFuel := -1
	for pos := crabs[0]; pos <= crabs[len(crabs)-1]; pos++ {
		var fuel int
		for _, p := range crabs {
			moves := int(math.Abs(float64(pos - p)))
			switch moves {
			case 1:
				fuel++
				fallthrough
			case 0:
				continue
			}
			fuel += int(float64(moves+1) * float64(moves) / 2)
		}
		if minFuel < 0 || fuel < minFuel {
			minFuel = fuel
		}
	}

	return minFuel
}
