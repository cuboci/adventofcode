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
		log.Fatal("missing input")
	}

	codes, err := read(flag.Arg(0))
	if err != nil {
		log.Fatal(err)
	}

	power, err := power(codes)
	if err != nil {
		log.Fatal(err)
	}
	life, err := lifesupport(codes)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("power consumption: %d\n", power)
	fmt.Printf("life support: %d\n", life)
}

func read(file string) ([]string, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var codes []string
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		codes = append(codes, sc.Text())
	}
	if err := sc.Err(); err != nil {
		return nil, err
	}

	return codes, nil
}

func filter(codes []string, pos int) ([]string, []string, error) {
	var ones, zeros []string

	for _, code := range codes {
		if pos < 0 || pos >= len(code) {
			return nil, nil, fmt.Errorf("invalid position: %d", pos)
		}
		switch code[pos] {
		case '0':
			zeros = append(zeros, code)
		case '1':
			ones = append(ones, code)
		default:
			return nil, nil, fmt.Errorf("invalid code: %s", code)
		}
	}

	return zeros, ones, nil
}

func power(codes []string) (uint32, error) {
	var (
		gamma, epsilon uint32
		codeLen        = len(codes[0])
	)

	for pos := 0; pos < codeLen; pos++ {
		gamma <<= 1
		epsilon <<= 1
		epsilon |= 1
		zeros, ones, err := filter(codes, pos)
		if err != nil {
			return 0, err
		}
		if len(ones) > len(zeros) {
			gamma |= 1
			epsilon &= 0xfffffffe
		}
	}

	return gamma * epsilon, nil
}

func lifesupport(codes []string) (uint32, error) {
	var (
		codeLen     = len(codes[0])
		filtered    = codes
		ones, zeros []string
		err         error
	)

	for pos := 0; pos < codeLen && len(filtered) > 1; pos++ {
		filtered, ones, err = filter(filtered, pos)
		if err != nil {
			return 0, err
		}
		if len(ones) >= len(filtered) {
			filtered = ones
		}
	}

	if len(filtered) != 1 {
		return 0, fmt.Errorf("could not find oxygen rate")
	}

	oxygen, err := strconv.ParseUint(filtered[0], 2, codeLen)
	if err != nil {
		return 0, fmt.Errorf("invalid code: %w", err)
	}

	filtered = codes
	for pos := 0; pos < codeLen && len(filtered) > 1; pos++ {
		zeros, filtered, err = filter(filtered, pos)
		if err != nil {
			return 0, err
		}
		if len(zeros) <= len(filtered) {
			filtered = zeros
		}
	}

	if len(filtered) != 1 {
		return 0, fmt.Errorf("could not find CO₂ scrubber rate")
	}

	co2, err := strconv.ParseUint(filtered[0], 2, codeLen)
	if err != nil {
		return 0, fmt.Errorf("invalid code: %w", err)
	}

	return uint32(oxygen * co2), nil
}
