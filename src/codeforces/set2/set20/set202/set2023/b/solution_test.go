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
	s := `2
15 16
2 1`
	expect := 16
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `5
10 10 100 100 1000
3 4 1 1 1`
	expect := 200
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `3
100 49 50
3 2 2`
	expect := 100
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `4
100 200 300 1000
2 3 4 1`
	expect := 1000
	runSample(t, s, expect)
}
