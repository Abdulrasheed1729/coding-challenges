package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	lines := flag.Bool("l", false, "Count lines")

	bytes := flag.Bool("b", false, "Count bytes")

	chars := flag.Bool("c", false, "Count characters")

	flag.Parse()

	flagArgs := flag.Args()
	var file string

	if flagArgs == nil {
		fmt.Println(count(os.Stdin, *lines, *bytes, *chars))
	}
	file = flagArgs[0]
	f, err := os.Open(file)

	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	fmt.Println(count(f, *lines, *bytes, *chars))

}

func count(r io.Reader, countLines, countBytes, countChars bool) int {
	scanner := bufio.NewScanner(r)

	if !countLines {
		scanner.Split(bufio.ScanWords)
	}

	if countBytes {
		scanner.Split(bufio.ScanBytes)
	}

	if countChars {
		scanner.Split(bufio.ScanRunes)
	}

	wc := 0

	for scanner.Scan() {
		wc++
	}

	return wc
}
