package main

import "testing"

func runSample(t *testing.T, board []string, expect int) {
	res := solve(board)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	board := []string{
		"01",
		"11",
	}
	expect := 2
	runSample(t, board, expect)
}

func TestSample2(t *testing.T) {
	board := []string{
		"01",
		"01",
	}
	expect := 2
	runSample(t, board, expect)
}

func TestSample3(t *testing.T) {
	board := []string{
		"0101001010",
		"1010100110",
	}
	expect := 6
	runSample(t, board, expect)
}
