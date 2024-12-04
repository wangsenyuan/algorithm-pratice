package main

import "testing"

func runSample(t *testing.T, s string, expect string) {
	res := solve(s)

	x := countMinMoves(s, expect)
	y := countMinMoves(s, res)

	if x != y {
		t.Fatalf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "ANTON"
	expect := "NNOTA"
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "NAAN"
	expect := "AANN"
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := "AAAAAA"
	expect := "AAAAAA"
	runSample(t, s, expect)
}


func TestSample4(t *testing.T) {
	s := "OAANTTON"
	expect := "TNNTAOOA"
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := "NONOTA"
	expect := "ATOONN"
	runSample(t, s, expect)
}