package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type movement string

const (
	forward = movement("forward")
	up      = movement("up")
	down    = movement("down")
)

type command struct {
	dir  movement
	dist int
}

type submarine struct {
	pos, depth, aim int
}

func (s *submarine) move1(cmd command) {
	switch cmd.dir {
	case forward:
		s.pos += cmd.dist
	case down:
		s.depth += cmd.dist
	case up:
		s.depth -= cmd.dist
	}
}

func (s *submarine) move2(cmd command) {
	switch cmd.dir {
	case forward:
		s.pos += cmd.dist
		s.depth += s.aim * cmd.dist
	case down:
		s.aim += cmd.dist
	case up:
		s.aim -= cmd.dist
	}
}

func (s *submarine) position() int {
	return s.pos * s.depth
}

func main() {
	flag.Parse()
	if flag.NArg() < 1 {
		log.Fatal("missing input data")
	}

	commands, err := read(flag.Arg(0))
	if err != nil {
		log.Fatal(err)
	}

	sub1 := &submarine{}
	sub2 := &submarine{}

	for _, cmd := range commands {
		sub1.move1(cmd)
		sub2.move2(cmd)
	}

	fmt.Printf("position 1: %d\n", sub1.position())
	fmt.Printf("position 2: %d\n", sub2.position())
}

func read(file string) ([]command, error) {
	var res []command
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	for sc.Scan() {
		line := sc.Text()
		cmd, err := parseCommand(line)
		if err != nil {
			return nil, err
		}
		res = append(res, cmd)
	}
	if err := sc.Err(); err != nil {
		return nil, err
	}

	return res, nil
}

func parseCommand(val string) (command, error) {
	fields := strings.Fields(val)
	if len(fields) != 2 {
		return command{}, fmt.Errorf("invalid command: %s", val)
	}

	dist, err := strconv.Atoi(fields[1])
	if err != nil {
		return command{}, fmt.Errorf("invalid distance: %s: %w", fields[1], err)
	}

	cmd := command{dist: dist}
	switch fields[0] {
	case "forward":
		cmd.dir = forward
	case "down":
		cmd.dir = down
	case "up":
		cmd.dir = up
	default:
		return command{}, fmt.Errorf("invalid movement: %s", fields[0])
	}

	return cmd, nil
}
