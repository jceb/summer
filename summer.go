package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	goopt "github.com/droundy/goopt"
)

const VERSION = "0.1"

type Opts struct {
	prnt bool
	delimiter string
	field int
}

func SumLine(line string, opts *Opts) float64 {
	var field string
	var fields []string

	if opts.delimiter == "" {
		// take any space as separator, like awk
		fields = strings.Fields(line)
	} else {
		// a specific separator has been specified
		fields = strings.Split(line, opts.delimiter)
	}

	if len(fields) > opts.field {
		fields = strings.Fields(fields[opts.field])
		if len(fields) > 0 {
			field = fields[0]
		} else {
			return 0
		}
	} else {
		return 0
	}

	value, e := strconv.ParseFloat(field, 32)
	if e == strconv.ErrSyntax {
		return 0
	}

	return value
}

func SumString(s string, opts *Opts) (float64, string) {
	remainder := ""
	// FIXME make newline cross platform compatible
	idx := strings.Index(s, "\n")
	len_l := len(s)
	offset := 0
	var sum float64 = 0

	// compute value for field in string
	for idx != -1 && offset < len_l {
		if opts.prnt {
			fmt.Println(s[offset:offset+idx])
		}
		sum += SumLine(s[offset:offset+idx], opts)

		// increase offset and idx
		offset += idx + 1
		if offset < len_l {
			// FIXME make newline cross platform compatible
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
	var opts *Opts = new(Opts)
	stream := make([]byte, 1024)

	f := goopt.Int([]string{"-f", "--field"}, 1, "Selected field")
	d := goopt.String([]string{"-d", "--delimiter"}, "", "Use delimiter instead of space-like characters")
	p := goopt.Flag([]string{"-n", "--no-print"}, []string{"-p", "--print"}, "Don't print input", "Print input")
	goopt.Version = VERSION
	goopt.Summary = "Sum values in selected field"
	goopt.Parse(nil)

	opts.field = *f
	opts.delimiter = *d
	opts.prnt = !(*p)

	// start counting fields at 0
	if int(opts.field) < 1 {
		fmt.Fprintln(os.Stderr, "ERROR: Field must be bigger than 1")
		os.Exit(1)
	}
	opts.field -= 1


	// read input from STDIN
	for n, err := os.Stdin.Read(stream); n > 0 && err == nil; n, err = os.Stdin.Read(stream) {
		s := string(stream[:n])
		// join remainder and s
		if remainder != "" {
			s = strings.Join([]string{remainder, s}, "")
			// truncate remainder since it's part of s now
			remainder = ""
		}
		res, remainder = SumString(s, opts)
		sum += res
	}

	// print sum
	fmt.Println(sum)
}
