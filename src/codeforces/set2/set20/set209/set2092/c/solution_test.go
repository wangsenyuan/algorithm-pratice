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
	s := `3
5 3 9`
	expect := 9
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `2
3 2`
	expect := 5
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `4
1 2 2 1`
	expect := 5
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `5
5 4 3 2 9`
	expect := 21
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := `4
7 9 4 11`
// 7, 12, 1, 11
// 18, 1, 1, 11
	expect := 29
	runSample(t, s, expect)
}
