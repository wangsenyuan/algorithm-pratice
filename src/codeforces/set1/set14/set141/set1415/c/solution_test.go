package main

import "testing"

func runSample(t *testing.T, p int, k int, s string, x int, y int, expect int) {
	res := solve(p, k, s, x, y)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	p, k := 3, 2
	s := "0101010101"
	x, y := 2, 2
	expect := 2
	runSample(t, p, k, s, x, y, expect)
}

func TestSample2(t *testing.T) {
	p, k := 4, 1
	s := "00000"
	x, y := 2, 10
	expect := 4
	runSample(t, p, k, s, x, y, expect)
}

func TestSample3(t *testing.T) {
	p, k := 2, 3
	s := "10110011000"
	x, y := 4, 3
	expect := 10
	runSample(t, p, k, s, x, y, expect)
}
