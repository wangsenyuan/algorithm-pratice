package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect string) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := process(reader)
	reader = bufio.NewReader(strings.NewReader(expect))
	for _, x := range res {
		y := readString(reader)
		if x != y {
			t.Fatalf("Sample expect %s, but got %v", expect, res)
		}
	}
}

func TestSample1(t *testing.T) {
	s := `2 3
.*
..
`
	expect := `.*******
..******
.*.*****
....****
.***.***
..**..**
.*.*.*.*
........
`
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `3 2
.*.
***
.*.
`
	expect := `.*.***.*.
*********
.*.***.*.
*********
*********
*********
.*.***.*.
*********
.*.***.*.
`
	runSample(t, s, expect)
}
