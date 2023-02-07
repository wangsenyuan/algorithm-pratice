package main

import "testing"

func runSample(t *testing.T, x int, y int, a string, b string, expect int) {
	res := int(solve(x, y, a, b))

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	x, y := 3, 4
	a := "010001"
	b := "101000"
	expect := 7
	runSample(t, x, y, a, b, expect)
}

func TestSample2(t *testing.T) {
	x, y := 8, 3
	a := "0111001"
	b := "0100001"
	expect := solve1(x, y, a, b)
	runSample(t, x, y, a, b, int(expect))
}
