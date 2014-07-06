package main

import (
	"fmt"
	"os"
	"math"
	"strconv"
	"strings"
	"time"
	goopt "github.com/droundy/goopt"
)

const VERSION = "0.1"

type Sum struct {
	f float64
	d time.Duration
}

type Opts struct {
	prnt bool
	delimiter string
	field int
	sum_type int // type -1 undecided, 1 float, 2 duration
}

func SumLine(line string, opts *Opts, sum *Sum) {
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
			return
		}
	} else {
		return
	}

	// parse duration
	if opts.sum_type == 2 || opts.sum_type == -1 {
		_d, e := time.ParseDuration(field)
		if e == nil {
			sum.d += _d
			if opts.sum_type == -1 {
				opts.sum_type = 2
			}
		}
	}
	// parse float
	if opts.sum_type == 1 || opts.sum_type == -1 {
		_f, e := strconv.ParseFloat(field, 32)
		if e == nil {
			sum.f += _f
			if opts.sum_type == -1 {
				opts.sum_type = 1
			}
		}
	}
}

func SumString(s string, opts *Opts, sum *Sum) string {
	remainder := ""
	idx := strings.Index(s, "\n")
	len_l := len(s)
	offset := 0

	// compute value for field in string
	for idx != -1 && offset < len_l {
		if opts.prnt {
			fmt.Println(s[offset:offset+idx])
		}
		SumLine(s[offset:offset+idx], opts, sum)

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
	return remainder
}

func Round(value float64, digits int) float64 {
	if digits >= 0 {
		scale := math.Pow(10, float64(digits))
		return float64(int(math.Floor((value * scale)+0.5))) / scale
	}
	return value
}

func main() {
	var sum Sum = Sum{0, time.Duration(0)}
	var remainder string
	var opts Opts = Opts{true, "", 1, -1}
	stream := make([]byte, 1024)

	f := goopt.Int([]string{"-f", "--field"}, opts.field, "Selected field")
	d := goopt.String([]string{"-d", "--delimiter"}, opts.delimiter, "Use delimiter instead of space-like characters")
	p := goopt.Flag([]string{"-n", "--no-print"}, []string{"-p", "--print"}, "Don't print input", "Print input")
	s := goopt.Int([]string{"-s", "--scale"}, 2, "Scale to number of digits after the decimal point")
	goopt.Version = VERSION
	goopt.Summary = "Sum values in column and print results"
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
		remainder = SumString(input, &opts, &sum)
	}
	if opts.sum_type == 1 {
		fmt.Printf("%." + strconv.Itoa(*s) + "f\n", Round(sum.f, *s))
	} else {
		fmt.Println(sum.d)
	}
}
