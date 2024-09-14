package main

import (
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	res, color := solve(s)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d %v", expect, res, color)
	}
}

func TestSample1(t *testing.T) {
	s := "abacbecfd"
	expect := 2
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "aaabbcbb"
	expect := 2
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := "raaaaabbbccccccccccdddddeeeeeeeeeeeeeeffffffffffggggggggghhhhhiiiiiiiiijjjjkkkkkkkkkkllllllllmmmmmmmmnnnnnnnooooooooppppppqqqqqqqqqqrrrrrrrrssssttttttttttuuuuuuuuvvvvvvwwwwwwxxxxxxxyyyyyyyzzzzzzzzzzzz"
	expect := 2
	runSample(t, s, expect)
}
