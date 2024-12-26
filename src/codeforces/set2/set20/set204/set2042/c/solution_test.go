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
	s := `4 1
1001`
	expect := 2
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `4 1
1010`
	expect := -1
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `4 1
0110`
	expect := 2
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `4 2
0110`
	expect := -1
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := `6 3
001110`
	expect := 3
	runSample(t, s, expect)
}

func TestSample6(t *testing.T) {
	s := `10 20
1111111111`
	expect := 4
	runSample(t, s, expect)
}

func TestSample7(t *testing.T) {
	s := `5 11
11111`
	expect := -1
	runSample(t, s, expect)
}
