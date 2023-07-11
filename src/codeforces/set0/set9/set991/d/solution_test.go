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
		"00",
		"00",
	}
	expect := 1
	runSample(t, board, expect)
}

func TestSample2(t *testing.T) {
	board := []string{
		"00X00X0XXX0",
		"0XXX0X00X00",
	}
	expect := 4
	runSample(t, board, expect)
}

func TestSample3(t *testing.T) {
	board := []string{
		"0X0X0",
		"0X0X0",
	}
	expect := 0
	runSample(t, board, expect)
}
