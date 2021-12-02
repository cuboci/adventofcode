package main

import (
	"crypto/md5"
	"flag"
	"fmt"
	"io"
	"log"
	"strings"
)

func main() {
	flag.Parse()
	if flag.NArg() < 1 {
		log.Fatal("missing input data")
	}

	fmt.Printf("coin: %d\n", mine(flag.Arg(0), 5))
	fmt.Printf("coin: %d\n", mine(flag.Arg(0), 6))
}

func mine(input string, prefLen int) int {
	prefix := fmt.Sprintf(fmt.Sprintf("%%0%ds", prefLen), "")
	hash := md5.New()
	for i := 1; ; i++ {
		hash.Reset()
		io.WriteString(hash, fmt.Sprintf("%s%d", input, i))
		if strings.HasPrefix(fmt.Sprintf("%16X", hash.Sum(nil)), prefix) {
			return i
		}
	}
}
