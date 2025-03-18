package main

import (
	"bufio"
	"math"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect float64) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := process(reader)
	if math.Abs(res-expect) > eps {
		t.Errorf("Sample expect %f, but got %f", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `3
-5 9
0 1
5 -1
`
	expect := 1.0
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `6
1 3
2 3
3 3
4 -3
5 -1
6 -100
`
	expect := 0.02912621359223301065

	runSample(t, s, expect)
}
