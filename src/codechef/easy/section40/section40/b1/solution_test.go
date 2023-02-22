package main

import "testing"

func runSample(t *testing.T, board []string, expect int) {
	res := solve(board)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	board := []string{
		"....$",
		".....",
		".....",
		"D...D",
		".....",
		".....",
		"@....",
	}
	expect := 2
	runSample(t, board, expect)
}

func TestSample2(t *testing.T) {
	board := []string{
		"....$",
		".....",
		".....",
		"DDDDD",
		".....",
		".....",
		"@....",
	}
	expect := 0
	runSample(t, board, expect)
}
