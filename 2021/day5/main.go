package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	flag.Parse()
	if flag.NArg() < 1 {
		log.Fatal("missing input")
	}

	lines, maxX, maxY, err := read(flag.Arg(0))
	if err != nil {
		log.Fatal(err)
	}

	oceanFloor := make([]int, (maxX+1)*(maxY+1))
	for _, line := range lines {
		switch {
		case line.horizontal():
			if line.x1 > line.x2 {
				line.x1, line.x2 = line.x2, line.x1
			}
			for i := line.x1; i <= line.x2; i++ {
				val := (maxX+1)*line.y1 + i
				oceanFloor[val]++
			}
		case line.vertical():
			if line.y1 > line.y2 {
				line.y1, line.y2 = line.y2, line.y1
			}
			for i := line.y1; i <= line.y2; i++ {
				val := i*(maxX+1) + line.x1
				oceanFloor[val]++
			}
		}
	}

	var overlap int
	for _, val := range oceanFloor {
		if val > 1 {
			overlap++
		}
	}

	fmt.Printf("intersections: %d\n", overlap)

	for _, line := range lines {
		if line.horizontal() || line.vertical() {
			continue
		}
		div := line.x2 - line.x1
		if line.x1 > line.x2 {
			div = line.x1 - line.x2
		}
		start := line.y1*(maxX+1) + line.x1
		end := line.y2*(maxX+1) + line.x2
		if start > end {
			start, end = end, start
		}
		for i := start; i <= end; i += (end - start) / div {
			oceanFloor[i]++
		}
	}

	overlap = 0
	for _, val := range oceanFloor {
		if val > 1 {
			overlap++
		}
	}

	fmt.Printf("intersections: %d\n", overlap)
}

func read(file string) ([]line, int, int, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, 0, 0, err
	}
	defer f.Close()

	var (
		lines      []line
		line       line
		maxX, maxY int
		rx         = regexp.MustCompile(`^(\d+),(\d+) -> (\d+),(\d+)$`)
	)

	sc := bufio.NewScanner(f)
	for sc.Scan() {
		l := sc.Text()
		m := rx.FindStringSubmatch(l)
		line.x1, _ = strconv.Atoi(m[1])
		line.y1, _ = strconv.Atoi(m[2])
		line.x2, _ = strconv.Atoi(m[3])
		line.y2, _ = strconv.Atoi(m[4])

		if line.x1 > maxX {
			maxX = line.x1
		}
		if line.x2 > maxX {
			maxX = line.x2
		}
		if line.y1 > maxY {
			maxY = line.y1
		}
		if line.y2 > maxY {
			maxY = line.y2
		}
		lines = append(lines, line)
	}
	if err := sc.Err(); err != nil {
		return nil, 0, 0, err
	}

	return lines, maxX, maxY, nil
}

type line struct {
	x1, x2, y1, y2 int
}

func (l line) String() string {
	return fmt.Sprintf("%d,%d -> %d,%d", l.x1, l.y1, l.x2, l.y2)
}

func (l line) horizontal() bool {
	return l.y1 == l.y2
}

func (l line) vertical() bool {
	return l.x1 == l.x2
}
