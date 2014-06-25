package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

// func parse_string(s string) (int, string) {
// }

func main() {
	// sum
	var sum float64 = 0

	var prnt bool
	flag.BoolVar(&prnt, "p", false, "Print input")
	del := flag.String("d", "\t", "Use delimiter instead of tab")
	del_isspace := unicode.IsSpace(rune((*del)[0]))
	var field int
	flag.IntVar(&field, "f", 1, "Selected field")

	flag.Parse()

	if field < 1 {
		fmt.Fprintln(os.Stderr, "ERROR: Field must be bigger than 1")
		os.Exit(1)
	}
	// start counting fields at 0
	field -= 1

	var remainder string
	stream := make([]byte, 1024)
	// read stream
	for n, err := os.Stdin.Read(stream); n > 0 && err == nil; n, err = os.Stdin.Read(stream) {
		s := string(stream[:n])
		// join remainder and s
		if remainder != "" {
			s = strings.Join([]string{remainder, s}, "")
			// truncate remainder since it's part of s now
			remainder = ""
		}
		idx := strings.Index(s, "\n")
		len_s := len(s)
		offset := 0

		// compute value for field in line(s)
		for idx != -1 && offset < len_s {
			if prnt {
				fmt.Println(s[offset:offset+idx])
			}
			var _field string
			if del_isspace {
				// take any space as separator, like awk
				f := strings.Fields(s[offset:offset+idx])
				if len(f) > field {
					_field = f[field]
				} else {
					// ignore lines with not enough fields

					// FIXME code duplication
					offset += idx + 1
					if offset < len_s {
						idx = strings.Index(s[offset:], "\n")
					} else {
						break
					}
					continue
				}
			} else {
				// a specific separator has been specified
				strings.Split(*del, s[offset:offset+idx])
			}

			r, e := strconv.ParseFloat(_field, 32)
			if e == strconv.ErrSyntax {
				continue
			}
			sum += r

			// increase offset and idx
			offset += idx + 1
			if offset < len_s {
				idx = strings.Index(s[offset:], "\n")
			} else {
				break
			}
		}

		// extend remainder by new input
		if offset < len_s {
			remainder = strings.Join([]string{remainder, s[offset:len_s]}, "")
		}
	}

	// print sum
	fmt.Println(sum)
}
