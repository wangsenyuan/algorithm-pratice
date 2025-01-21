package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect bool) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := process(reader)
	if res != expect {
		t.Errorf("Sample %s, expect %t, but got %t", s, expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `6
3 6 9 18 36 108
`
	expect := true
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `2
7 17
`
	expect := false
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `9
4 8 10 12 15 18 33 44 81
`
	expect := true
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `7
7 38 690 717 11165 75361 104117
`
	expect := true
	runSample(t, s, expect)
}
