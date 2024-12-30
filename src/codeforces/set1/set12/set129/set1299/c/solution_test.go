package main

import (
	"bufio"
	"math"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect []float64) {
	reader := bufio.NewReader(strings.NewReader(s))

	res := process(reader)

	for i := range len(res) {
		if math.Abs(res[i]-expect[i]) > eps {
			t.Fatalf("Sample expect %v, but got %v", expect, res)
		}
	}
}

func TestSample1(t *testing.T) {
	s := `4
7 5 5 7
`
	expect := []float64{
		5.666666667,
		5.666666667,
		5.666666667,
		7.000000000,
	}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `5
7 8 8 10 12
`
	expect := []float64{
		7.000000000,
		8.000000000,
		8.000000000,
		10.000000000,
		12.000000000,
	}
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `10
3 9 5 5 1 7 5 3 8 7
`
	expect := []float64{
		3.000000000,
		5.000000000,
		5.000000000,
		5.000000000,
		5.000000000,
		5.000000000,
		5.000000000,
		5.000000000,
		7.500000000,
		7.500000000,
	}
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `12
8 10 4 6 6 4 1 2 2 6 9 5
`
	expect := []float64{
		4.777777778,
		4.777777778,
		4.777777778,
		4.777777778,
		4.777777778,
		4.777777778,
		4.777777778,
		4.777777778,
		4.777777778,
		6.000000000,
		7.000000000,
		7.000000000,
	}
	runSample(t, s, expect)
}
