package main

import "testing"

func runSample(t *testing.T, x, y int, expect int) {
	res := solve(x, y)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	x, y := 2, 7
	expect := 1
	runSample(t, x, y, expect)
}


func TestSample2(t *testing.T) {
	x, y := 1, 5
	expect := 2
	runSample(t, x, y, expect)
}

func TestSample3(t *testing.T) {
	x, y := 3, 10
	expect := 1
	runSample(t, x, y, expect)
}
