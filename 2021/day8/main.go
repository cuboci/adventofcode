package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	flag.Parse()
	if flag.NArg() < 1 {
		log.Fatal("missing input")
	}

	displays, err := read(flag.Arg(0))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("number of unique digits (1, 4, 7, 8): %d\n", countUniqueDigits(displays))
	fmt.Printf("sum of display values: %d\n", addedDisplayValues(displays))
}

func read(file string) ([]display, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var ret []display
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		d, err := newDisplay(sc.Text())
		if err != nil {
			return nil, err
		}
		ret = append(ret, d)
	}
	if err := sc.Err(); err != nil {
		return nil, err
	}

	return ret, nil
}

type wire int

const (
	wto wire = 1 << iota
	wul
	wur
	wmi
	wll
	wlr
	wbo
)

var digitValues = map[wire]int{
	wto + wul + wur + wll + wlr + wbo:       0,
	wur + wlr:                               1,
	wto + wur + wmi + wll + wbo:             2,
	wto + wur + wmi + wlr + wbo:             3,
	wul + wur + wmi + wlr:                   4,
	wto + wul + wmi + wlr + wbo:             5,
	wto + wul + wmi + wll + wlr + wbo:       6,
	wto + wur + wlr:                         7,
	wto + wul + wur + wmi + wll + wlr + wbo: 8,
	wto + wul + wur + wmi + wlr + wbo:       9,
}

var wireSignatures = map[int]wire{
	51: wto,
	40: wul,
	46: wur,
	45: wmi,
	28: wll,
	53: wlr,
	47: wbo,
}

func countUniqueDigits(displays []display) int {
	var ret int
	for _, d := range displays {
		for _, code := range d.code {
			switch len(code) {
			case 2, 3, 4, 7:
				ret++
			}
		}
	}

	return ret
}

func addedDisplayValues(displays []display) int {
	var ret int
	for _, d := range displays {
		ret += d.value()
	}

	return ret
}
