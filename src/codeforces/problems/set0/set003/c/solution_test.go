package main

import "testing"

func runSample(t *testing.T, board []string, expect string) {
	res := solve(board)

	if res != expect {
		t.Errorf("Sample %v, expect %s, but got %s", board, expect, res)
	}
}

func TestSample1(t *testing.T) {
	board := []string{
		"X0X",
		".0.",
		".X.",
	}
	expect := "second"
	runSample(t, board, expect)
}
