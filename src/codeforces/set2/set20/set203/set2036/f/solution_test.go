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
	s := "1 3 1 0"
	expect := 2
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "2 28 3 7"
	expect := 2
	// 111, 1111
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := "15 43 1 0"
	expect := 13
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := "1 1000000000 30 1543"
	expect := 1000000519
	runSample(t, s, expect)
}
