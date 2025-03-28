package main

import (
	"bufio"
	"math"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect float64) {
	reader := bufio.NewReader(strings.NewReader(s))
	ans := process(reader)
	if math.Abs(ans-expect) > 1e-10 {
		t.Errorf("Sample expect %f, but got %f", expect, ans)
	}
}

func TestSample1(t *testing.T) {
	s := `1
1 2
50
`
	expect := 0.5
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `2
1 2
9 11
50
`
	expect := 0.833333333333333
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `1
1 1000000000000000000
50
`
	expect := 0.111111111111111
	runSample(t, s, expect)
}
