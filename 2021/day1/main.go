package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	flag.Parse()

	if flag.NArg() < 1 {
		log.Fatal("missing input data")
	}

	data, err := read(flag.Arg(0))
	if err != nil {
		log.Fatal(err)
	}

	increased := simple(data)
	fmt.Println(increased)
	increased = sliding(data)
	fmt.Println(increased)
}

func read(datafile string) ([]int, error) {
	f, err := os.Open(datafile)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	var res []int

	for sc.Scan() {
		val, err := strconv.Atoi(sc.Text())
		if err != nil {
			return nil, err
		}
		res = append(res, val)
	}
	if err := sc.Err(); err != nil {
		return nil, err
	}

	return res, nil
}

func simple(data []int) int {
	var depth, increased int
	for _, val := range data {
		if val > depth {
			if depth > 0 {
				increased++
			}
		}
		depth = val
	}

	return increased
}

func sliding(data []int) int {
	res := make([]int, len(data)-2)
	for i := 0; i < (len(data) - 2); i++ {
		res[i] += data[i] + data[i+1] + data[i+2]
	}
	return simple(res)
}
