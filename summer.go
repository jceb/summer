package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	// "unicode"
)

func parse_line(line string, delimiter string, field int) float64 {
	var _field string

	if delimiter == "" {
		// take any space as separator, like awk
		fields := strings.Fields(line)
		if len(fields) > field {
			_field = fields[field]
		} else {
			// ignore lines with not enough fields
			return 0
		}
	} else {
		// a specific separator has been specified
		strings.Split(delimiter, line)
	}

	value, e := strconv.ParseFloat(_field, 32)
	if e == strconv.ErrSyntax {
		return 0
	}

	return value
}

func parse_string(s string, delimiter string, field int, prnt bool) (float64, string) {
	remainder := ""
	idx := strings.Index(s, "\n")
	len_l := len(s)
	offset := 0
	var sum float64 = 0

	// compute value for field in string
	for idx != -1 && offset < len_l {
		if prnt {
			fmt.Println(s[offset:offset+idx])
		}
		sum += parse_line(s[offset:offset+idx], delimiter, field)

		// increase offset and idx
		offset += idx + 1
		if offset < len_l {
			idx = strings.Index(s[offset:], "\n")
		} else {
			break
		}

		// extend remainder by new input
		if offset < len_l {
			remainder = strings.Join([]string{remainder, s[offset:len_l]}, "")
		}
	}

	return sum, remainder
}

func main() {
	// sum
	var sum float64 = 0
	var res float64
	var remainder string
	stream := make([]byte, 1024)

	prnt := flag.Bool("p", false, "Print input")
	del := flag.String("d", "", "Use delimiter instead of space-like characters")
	// del_isspace := unicode.IsSpace(rune((*del)[0]))
	field := flag.Int("f", 1, "Selected field")
	flag.Parse()

	// start counting fields at 0
	if *field < 1 {
		fmt.Fprintln(os.Stderr, "ERROR: Field must be bigger than 1")
		os.Exit(1)
	}
	*field -= 1


	// read input from STDIN
	for n, err := os.Stdin.Read(stream); n > 0 && err == nil; n, err = os.Stdin.Read(stream) {
		s := string(stream[:n])
		// join remainder and s
		if remainder != "" {
			s = strings.Join([]string{remainder, s}, "")
			// truncate remainder since it's part of s now
			remainder = ""
		}
		res, remainder = parse_string(s, *del, *field, *prnt)
		sum += res
	}

	// print sum
	fmt.Println(sum)
}
