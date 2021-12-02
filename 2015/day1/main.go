package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	flag.Parse()
	if flag.NArg() < 1 {
		log.Fatal("missing input file")
	}

	input, err := os.ReadFile(flag.Arg(0))
	if err != nil {
		log.Fatal(err)
	}
	input = bytes.TrimSpace(input)

	var floor int
	basementEntered := -1
	for i, dir := range string(input) {
		switch dir {
		case '(':
			floor++
		case ')':
			floor--
		default:
			log.Printf("dir: %q", dir)
			log.Fatal("Santa doesn't know where to go")
		}
		if basementEntered < 0 {
			if floor < 0 {
				basementEntered = i + 1
			}
		}
	}

	fmt.Printf("floor: %d\n", floor)
	fmt.Printf("basement entered at position: %d\n", basementEntered)
}
