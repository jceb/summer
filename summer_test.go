package main

import (
	"testing"
	"time"
)

func SumFloat(t *testing.T, in *string, out *[]float64, digits int){
	for i := range *out {
		var sum Sum = Sum{0, time.Duration(0)}
		var opts Opts = Opts{false, ";", i, -1}
		if SumString(*in, &opts, &sum); Round(sum.f, digits) != (*out)[i] {
			t.Errorf("SumString(%v) = %v, want %v", *in, Round(sum.f, digits), (*out)[i])
		}
	}
}

func SumDuration(t *testing.T, in *string, out *[]time.Duration){
	for i := range *out {
		var sum Sum = Sum{0, time.Duration(0)}
		var opts Opts = Opts{false, ";", i, -1}
		if SumString(*in, &opts, &sum); sum.d != (*out)[i] {
			t.Errorf("SumString(%v) = %v, want %v", *in, sum.d, (*out)[i])
		}
	}
}


func TestSimpleSum(t *testing.T) {
	var in string = "1;2\n3;4\n"
	var out []float64 = []float64{4, 6}

	SumFloat(t, &in, &out, -1)
}

func TestIgnoreNonNumbers(t *testing.T) {
	var in string = "1;2\na;1\n\n ; \n3;4\n"
	var out []float64 = []float64{4, 7}

	SumFloat(t, &in, &out, -1)
}

func TestIgnoreSpace(t *testing.T) {
	var in string = " \t1 ;2\t\na;   1   \n\t\t3;\t4\n"
	var out []float64 = []float64{4, 7}

	SumFloat(t, &in, &out, -1)
}

func TestSumNegativeValues(t *testing.T) {
	var in string = "1;2\n3;4\n-3;-9\n"
	var out []float64 = []float64{1, -3}

	SumFloat(t, &in, &out, -1)
}

func TestSumFloats(t *testing.T) {
	var in string = "0.1;0.2\n3.3;4.2\n-3;-9\n"
	var out []float64 = []float64{0.4, -4.6}

	SumFloat(t, &in, &out, 2)
}

const SECONDS = 1000000000

func TestSumDuration(t *testing.T) {
	var in string = "10h;11h\n3h;4h20m\n-3h;-9h\n"
	var out []time.Duration = []time.Duration{time.Duration(10*60*60*SECONDS), time.Duration(6*60*60*SECONDS + 20*60*SECONDS)}

	SumDuration(t, &in, &out)
}

func TestSumIsoTime(t *testing.T) {
	var in string = "10:00:01;13:00\n03:00;4:20\n3:00;9:00\n"
	var out []time.Duration = []time.Duration{time.Duration(16*60*60*SECONDS + 1*SECONDS), time.Duration(26*60*60*SECONDS + 20*60*SECONDS)}

	SumDuration(t, &in, &out)
}

func TestSumHugeIsoTime(t *testing.T) {
	var in string = "30:00:01;13:00\n03:70;4:20:3600\n3:00;9:00\n"
	var out []time.Duration = []time.Duration{time.Duration(37*60*60*SECONDS + 10*60*SECONDS + 1*SECONDS), time.Duration(27*60*60*SECONDS + 20*60*SECONDS)}

	SumDuration(t, &in, &out)
}

func TestSumNegativeIsoTime(t *testing.T) {
	var in string = "10:00:01;-11:00:00\n-;0:0\n-03:00;-4:20\n3:00;9:00\n"
	var out []time.Duration = []time.Duration{time.Duration(10*60*60*SECONDS + 1*SECONDS), time.Duration((6*60*60*SECONDS + 20*60*SECONDS) * -1)}

	SumDuration(t, &in, &out)
}
