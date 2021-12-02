package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"os"
)

type position struct {
	x, y int
}

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

	presents := doRun([]position{{}}, string(input))
	fmt.Printf("%d houses got at least one present\n", len(presents))
	presents = doRun([]position{{}, {}}, string(input))
	fmt.Printf("%d houses got at least one present a year later\n", len(presents))
}

func doRun(santas []position, input string) map[position]int {
	runs := len(santas)

	var santa position
	presents := make(map[position]int)

	for _, s := range santas {
		count := presents[s]
		presents[s] = count + 1
	}
	presents[santa] = 1
	for i, dir := range string(input) {
		santa = santas[i%runs]
		switch dir {
		case '>':
			santa.x++
		case '<':
			santa.x--
		case '^':
			santa.y++
		case 'v':
			santa.y--
		default:
			log.Fatal("Santa stepped into another dimension")
		}
		santas[i%runs] = santa
		count := presents[santa]
		presents[santa] = count + 1
	}

	return presents
}
