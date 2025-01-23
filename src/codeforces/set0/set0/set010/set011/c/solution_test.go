package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	res := process(bufio.NewReader(strings.NewReader(s)))

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `2 2
11
11
	`
	expect := 1
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `3 3
010
101
010
	`
	expect := 1
	runSample(t, s, expect)
}
func TestSample3(t *testing.T) {
	s := `8 8
00010001
00101000
01000100
10000010
01000100
00101000
11010011
11000011
	`
	expect := 1
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `10 10
1111111000
1000001000
1011001000
1011001010
1000001101
1001001010
1010101000
1001001000
1000001000
1111111000
	`
	expect := 2
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := `12 11
11111111111
10000000001
10111111101
10100000101
10101100101
10101100101
10100000101
10100000101
10111111101
10000000001
11111111111
00000000000
	`
	expect := 3
	runSample(t, s, expect)
}
