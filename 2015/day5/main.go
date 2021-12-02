package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	flag.Parse()
	if flag.NArg() < 1 {
		log.Fatal("missing input")
	}

	input, err := read(flag.Arg(0))
	if err != nil {
		log.Fatal(err)
	}

	var nice, nicer int
	for _, s := range input {
		if isNice(s) {
			nice++
		}
		if isNicer(s) {
			nicer++
		}
	}

	fmt.Printf("%d nice strings\n", nice)
	fmt.Printf("%d nicer strings\n", nicer)
}

func read(file string) ([]string, error) {
	var input []string

	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	for sc.Scan() {
		input = append(input, sc.Text())
	}
	if err := sc.Err(); err != nil {
		return nil, err
	}

	return input, nil
}

func isNice(s string) bool {
	var (
		vowels int
		double bool
		runes  = []rune(s)
	)

	for i, r := range runes {
		switch r {
		case 'a', 'e', 'i', 'o', 'u':
			vowels++
		}
		if i < len(runes)-1 {
			r1 := runes[i+1]
			if !double {
				double = r == r1
			}
			switch r {
			case 'a', 'c', 'p', 'x':
				if r1 == r+1 {
					return false
				}
			}
		}
	}

	return vowels >= 3 && double
}

func isNicer(s string) bool {
	var (
		double, between bool
		runes           = []rune(s)
	)
	for i := 0; i < len(runes)-2; i++ {
		double = double || (strings.Index(s[i+2:], string(runes[i:i+2])) >= 0)
		between = between || runes[i] == runes[i+2]
	}

	return double && between
}
