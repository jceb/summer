package main

import (
	"testing"
)

func TestSimpleSum(t *testing.T) {
	opts := new(Opts)
	opts.prnt, opts.delimiter, opts.field = false, ";", 1

	const in string = "1;2\n3;4\n"
	var out []float64 = []float64{4, 6}

	for i := range out {
		opts.field = i
		if sum, _ := SumString(in, opts); sum != out[i] {
			t.Errorf("SumString(%v) = %v, want %v", in, sum, out[i])
		}
	}
}

func TestIgnoreNonNumbers(t *testing.T) {
	opts := new(Opts)
	opts.prnt, opts.delimiter, opts.field = false, ";", 1

	const in string = "1;2\na;1\n\n ; \n3;4\n"
	var out []float64 = []float64{4, 7}

	for i := range out {
		opts.field = i
		if sum, _ := SumString(in, opts); sum != out[i] {
			t.Errorf("SumString(%v) = %v, want %v", in, sum, out[i])
		}
	}
}

func TestIgnoreSpace(t *testing.T) {
	opts := new(Opts)
	opts.prnt, opts.delimiter, opts.field = false, ";", 1

	const in string = " \t1 ;2\t\na;   1   \n\t\t3;\t4\n"
	var out []float64 = []float64{4, 7}

	for i := range out {
		opts.field = i
		if sum, _ := SumString(in, opts); sum != out[i] {
			t.Errorf("SumString(%v) = %v, want %v", in, sum, out[i])
		}
	}
}

func TestSumNegativeValues(t *testing.T) {
	opts := new(Opts)
	opts.prnt, opts.delimiter, opts.field = false, ";", 1

	const in string = "1;2\n3;4\n-3;-9\n"
	var out []float64 = []float64{1, -3}

	for i := range out {
		opts.field = i
		if sum, _ := SumString(in, opts); sum != out[i] {
			t.Errorf("SumString(%v) = %v, want %v", in, sum, out[i])
		}
	}
}

func TestSumFloats(t *testing.T) {
	opts := new(Opts)
	opts.prnt, opts.delimiter, opts.field = false, ";", 1

	const in string = "0.1;0.2\n3.3;4.2\n-3;-9\n"
	// var out []float64 = []float64{0.4, -4.6}
	var out []float64 = []float64{0.3999999538064003, -4.600000187754631}

	for i := range out {
		opts.field = i
		if sum, _ := SumString(in, opts); sum != out[i] {
			t.Errorf("SumString(%v) = %v, want %v", in, sum, out[i])
		}
	}
}
