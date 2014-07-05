package main

import (
	"fmt"
	"os"
	"math"
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

func Round(value float64, digits int) float64 {
	scale := math.Pow(10, float64(digits))
	return float64(int(math.Floor((value * scale)+0.5))) / scale
}

func main() {
	var sum float64 = 0
	var res float64
	var remainder string
	var opts *Opts = new(Opts)
	stream := make([]byte, 1024)

	f := goopt.Int([]string{"-f", "--field"}, 1, "Selected field")
	d := goopt.String([]string{"-d", "--delimiter"}, "", "Use delimiter instead of space-like characters")
	p := goopt.Flag([]string{"-n", "--no-print"}, []string{"-p", "--print"}, "Don't print input", "Print input")
	s := goopt.Int([]string{"-s", "--scale"}, 2, "Scale to number of digits after the decimal point")
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

	for n, err := os.Stdin.Read(stream); n > 0 && err == nil; n, err = os.Stdin.Read(stream) {
		// read input from STDIN
		input := string(stream[:n])
		if remainder != "" {
			// join remainder and input
			input = strings.Join([]string{remainder, input}, "")
			// truncate remainder since it's part of input now
			remainder = ""
		}
		res, remainder = SumString(input, opts)
		sum += res
	}
	fmt.Printf("%." + strconv.Itoa(*s) + "g\n", Round(sum, *s))
}
