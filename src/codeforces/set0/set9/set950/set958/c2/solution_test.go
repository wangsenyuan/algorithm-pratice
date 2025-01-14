package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	reader := bufio.NewReader(strings.NewReader(s))

	res := process(reader)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `4 3 10
3 4 7 2
`
	expect := 16
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `10 5 12
16 3 24 13 9 8 7 5 12 12
`
	expect := 37
	runSample(t, s, expect)
}
