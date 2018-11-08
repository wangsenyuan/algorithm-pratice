package main

import "testing"

func runSample(t *testing.T, n int, words []string, expect string) {
	res := solve(n, words)

	if res != expect {
		t.Errorf("Sample %d %v, expect %s, but got %s", n, words, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 6
	words := []string{
		"CODE",
		"TRAP",
		"ARCS",
		"CART",
		"ROAD",
		"PART",
	}
	expect := "CARTORORDCAAESDP"
	runSample(t, n, words, expect)
}

func TestSample2(t *testing.T) {
	n := 5
	words := []string{
		"ABCD",
		"EFGH",
		"IJKL",
		"MNOP",
		"QRST",
	}
	expect := "gridsnotpossible"
	runSample(t, n, words, expect)
}

func TestSample3(t *testing.T) {
	n := 1
	words := []string{
		"ABCD",
	}
	expect := "AAAAAAAAAAAAABCD"
	runSample(t, n, words, expect)
}
