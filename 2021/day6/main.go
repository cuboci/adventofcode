package main

import (
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

	population, err := read(flag.Arg(0))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("population after 80 days: %d\n", generate(population, 80))
	fmt.Printf("population after 256 days: %d\n", generate(population, 256))
}

func read(file string) (fishes, error) {
	data, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}

	tmp := strings.Split(strings.TrimSpace(string(data)), ",")
	population := make(fishes)
	for _, val := range tmp {
		d, err := strconv.Atoi(val)
		if err != nil {
			return nil, err
		}
		population[d]++
	}

	return population, nil
}

type fishes map[int]int64

func generate(population fishes, days int) int64 {
	for i := 0; i < days; i++ {
		x := make(fishes)
		for gen, c := range population {
			gen--
			if gen < 0 {
				x[8] = c
				gen = 6
			}
			x[gen] += c
		}
		population = x
	}

	var count int64
	for _, c := range population {
		count += c
	}

	return count
}
