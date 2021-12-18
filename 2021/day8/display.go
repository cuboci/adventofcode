package main

import (
	"fmt"
	"sort"
	"strings"
)

type display struct {
	patterns []string
	code     []string
	wires    map[string]wire
}

var rsort = func(data []rune) func(int, int) bool {
	return func(i, j int) bool {
		return data[i] < data[j]
	}
}

func newDisplay(input string) (display, error) {
	line := strings.Split(input, " | ")
	if len(line) != 2 {
		return display{}, fmt.Errorf("invalid input")
	}
	var d display
	for _, pat := range strings.Fields(line[0]) {
		rs := []rune(pat)
		sort.Slice(rs, rsort(rs))
		d.patterns = append(d.patterns, string(rs))
	}
	for _, code := range strings.Fields(line[1]) {
		rs := []rune(code)
		sort.Slice(rs, rsort(rs))
		d.code = append(d.code, string(rs))
	}
	d.decode()

	return d, nil
}

func (d *display) decode() {
	d.wires = make(map[string]wire)

	counts := make(map[string]int)
	for _, pat := range d.patterns {
		for _, c := range pat {
			counts[string(c)] += len(pat) + 1
		}
	}

	for c, count := range counts {
		d.wires[c] = wireSignatures[count]
	}
}

func (d display) value() int {
	var ret int
	for _, c := range d.code {
		ret *= 10
		var val wire
		for _, w := range c {
			val += d.wires[string(w)]
		}
		ret += digitValues[val]
	}

	return ret
}
