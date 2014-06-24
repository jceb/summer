package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// sum
	var sum float64 = 0

	del := flag.String("d", "\t", "Use delimiter instead of tab")
	var field int
	flag.IntVar(&field, "f", 1, "Selected field")

	flag.Parse()

	if field < 1 {
		fmt.Fprintln(os.Stderr, "ERROR: Field must be bigger than 1")
		os.Exit(1)
	}
	// start counting fields at 0
	field -= 1

	// initialize CSV reader
	reader := csv.NewReader(os.Stdin)
	reader.Comma = rune((*del)[0])
	reader.Comment = 0
	reader.FieldsPerRecord = -1
	reader.TrimLeadingSpace = true

	// read lines and sum fields
	for line, err := reader.Read(); err == nil; line, err = reader.Read() {
		r, e := strconv.ParseFloat(line[field], 32)
		if e == strconv.ErrSyntax {
			continue
		}
		sum += r
	}

	// print sum
	fmt.Println(sum)
}
