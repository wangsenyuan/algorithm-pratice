package main

import "testing"

func runSample(t *testing.T, n int, m int, x int, y int, expect bool) {
	res := solve(n, m, x, y)

	if res != expect {
		t.Errorf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, m, x, y := 2, 2, 2, 1
	expect := true
	runSample(t, n, m, x, y, expect)
}


func TestSample2(t *testing.T) {
	n, m, x, y := 2, 2, 1, 1
	expect := false
	runSample(t, n, m, x, y, expect)
}

func TestSample3(t *testing.T) {
	n, m, x, y := 2, 2, 1, 2
	expect := true
	runSample(t, n, m, x, y, expect)
}
