package main

import "testing"

func runSample(t *testing.T, x int, y int, expect int) {
	res := solve(x, y)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	x, y := 6, 3
	// 1, 1, 6
	// 1, 2, 3
	//

	expect := 36
	runSample(t, x, y, expect)
}

func TestSample2(t *testing.T) {
	x, y := 4, 2
	expect := 6
	runSample(t, x, y, expect)
}
