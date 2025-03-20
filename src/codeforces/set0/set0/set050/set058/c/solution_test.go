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
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `3
2 2 2
	`
	expect := 1
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `4
1 2 2 1
	`
	expect := 0
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `3
61452 50974 73849
	`
	expect := 2
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `10
1 1 2 3 4 4 3 2 1 10
	`
	expect := 9
	runSample(t, s, expect)
}
