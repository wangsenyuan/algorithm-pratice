package main

import "testing"

func runSample(t *testing.T, D int, S []string, r, c int, grid []string, expect string) {
	res := solve(D, S, r, c, grid)

	if res != expect {
		t.Errorf("Sample expect %s, but got %s", expect, res)
	}

}

func TestSample1(t *testing.T) {
	D := 5
	S := []string{"BED", "BAD", "RED", "BREED", "BEER"}
	r, c := 3, 3
	board := []string{
		"DER",
		"RAD",
		"FEE",
	}
	expect := "Bob"

	runSample(t, D, S, r, c, board, expect)
}

func TestSample2(t *testing.T) {
	D := 5
	S := []string{"BED", "BAD", "RED", "BREED", "BEER"}
	r, c := 3, 3
	board := []string{
		"BAR",
		"BEE",
		"RAN",
	}
	expect := "Alice"

	runSample(t, D, S, r, c, board, expect)
}
