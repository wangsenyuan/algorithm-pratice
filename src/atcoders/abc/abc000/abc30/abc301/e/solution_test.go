package main

import "testing"

func runSample(t *testing.T, board []string, k int, expect int) {
	res := solve(board, k)
	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	k := 5
	board := []string{
		"S.G",
		"o#o",
		".#.",
	}
	expect := 1
	runSample(t, board, k, expect)
}

func TestSample2(t *testing.T) {
	k := 1
	board := []string{
		"S.G",
		".#o",
		"o#.",
	}
	expect := -1
	runSample(t, board, k, expect)
}

func TestSample3(t *testing.T) {
	k := 2000000
	board := []string{
		"S.o..ooo..",
		"..o..o.o..",
		"..o..ooo..",
		"..o..o.o..",
		"..o..ooo.G",
	}
	expect := 18
	runSample(t, board, k, expect)
}
