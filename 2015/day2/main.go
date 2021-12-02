package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
)

type box struct {
	length, width, height int
}

func newBox(dimensions string) (box, error) {
	var b box
	_, err := fmt.Sscanf(dimensions, "%dx%dx%d", &b.length, &b.width, &b.height)
	if err != nil {
		return box{}, err
	}
	return b, nil
}

func (b box) area() int {
	dims := []int{b.length * b.width, b.width * b.height, b.height * b.length}

	var smallest, area int
	for _, dim := range dims {
		area += 2 * dim
		if smallest == 0 || dim < smallest {
			smallest = dim
		}
	}
	area += smallest

	return area
}

func (b box) ribbon() int {
	dims := []int{b.length, b.width, b.height}
	sort.Ints(dims)

	rl := dims[0]*2 + dims[1]*2 + b.length*b.width*b.height
	return rl
}

func main() {
	flag.Parse()
	if flag.NArg() < 1 {
		log.Fatal("missing input file")
	}

	boxes, err := read(flag.Arg(0))
	if err != nil {
		log.Fatal(err)
	}

	var paper, ribbon int
	for _, b := range boxes {
		paper += b.area()
		ribbon += b.ribbon()
	}

	fmt.Printf("total amount of wrapping paper: %d square feet\n", paper)
	fmt.Printf("total length of ribbon: %d feet\n", ribbon)
}

func read(file string) ([]box, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var boxes []box
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		b, err := newBox(sc.Text())
		if err != nil {
			return nil, fmt.Errorf("That's an oddly-shaped box: %w", err)
		}
		boxes = append(boxes, b)
	}
	if err := sc.Err(); err != nil {
		return nil, err
	}

	return boxes, nil
}
