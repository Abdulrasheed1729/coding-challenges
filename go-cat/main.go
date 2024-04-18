package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {

	numbers := flag.Bool("n", false, "Show line numbers.")

	flag.Parse()

	args := flag.Args()

	if len(args) == 0 {
		fmt.Print(cat(os.Stdin, *numbers))
	} else {

		for _, arg := range args {

			f, err := os.Open(arg)
			if err != nil {
				fmt.Println(err)
			}
			defer f.Close()
			fmt.Print(cat(f, *numbers))

		}
	}

}

func cat(r io.Reader, numbers bool) string {
	scanner := bufio.NewScanner(r)

	var stringBuffer strings.Builder

	line := 1

	for scanner.Scan() {
		var s string

		if numbers {
			s = fmt.Sprintf("%d  %s\n", line, scanner.Text())
		} else {

			s = fmt.Sprintf("%s\n", scanner.Text())
		}

		_, err := stringBuffer.WriteString(s)
		if err != nil {
			return fmt.Sprintln(err)
		}
		line++
	}

	result := stringBuffer.String()

	return result
}
