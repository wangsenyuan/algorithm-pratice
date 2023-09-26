package main

import "testing"

func runSample(t *testing.T, s string, x string, expect int) {
	res := solve(s, x)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "rggry"
	x := "r"
	expect := 3
	runSample(t, s, x, expect)
}

func TestSample2(t *testing.T) {
	s := "rrg"
	x := "r"
	expect := 2
	runSample(t, s, x, expect)
}

func TestSample3(t *testing.T) {
	s := "yrrgy"
	x := "y"
	expect := 4
	runSample(t, s, x, expect)
}

func TestSample4(t *testing.T) {
	s := "rgrgyrg"
	x := "r"
	expect := 1
	runSample(t, s, x, expect)
}

func TestSample5(t *testing.T) {
	s := "rrrgyyygy"
	x := "y"
	expect := 4
	runSample(t, s, x, expect)
}
