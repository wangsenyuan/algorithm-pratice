package main

import (
	"bufio"
	"math"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect [][]float64) {
	reader := bufio.NewReader(strings.NewReader(s))

	res := process(reader)

	for i := range len(expect) {
		a := expect[i]
		b := res[i]
		for j := range 2 {
			x := a[j]
			y := b[j]
			if math.Abs(y-x) > 1e-6 {
				t.Fatalf("Sample expect %v, but got %v", expect, res)
			}
		}
	}
}

func TestSample1(t *testing.T) {
	s := `2 10
0.7853
0.3
3
5.0 5.0
4.0 2.4
6.0 1.9
`
	expect := [][]float64{
		{5.000000000, 2.549499369},
		{4.000000000, 0.378324889},
	}

	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `2 10
0.7853
0.3
2
4.0 2.4
6.0 1.9
`
	expect := [][]float64{
		{10.204081436, 0.000000000},
		{4.000000000, 0.378324889},
	}

	runSample(t, s, expect)
}
