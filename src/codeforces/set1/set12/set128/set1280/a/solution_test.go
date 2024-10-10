package main

import "testing"

func runSample(t *testing.T, s string, x int, expect int) {
	res := solve(s, x)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "231"
	x := 5
	expect := 25
	runSample(t, s, x, expect)
}

func TestSample2(t *testing.T) {
	s := "2323"
	x := 7
	expect := 1438
	runSample(t, s, x, expect)
}

func TestSample3(t *testing.T) {
	s := "333"
	x := 6
	expect := 1101
	runSample(t, s, x, expect)
}

func TestSample4(t *testing.T) {
	s := "133321333"
	x := 24
	expect := 686531475
	runSample(t, s, x, expect)
}
