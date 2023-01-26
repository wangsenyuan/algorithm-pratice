package main

import "testing"

func runSample(t *testing.T, x, y int, A []int, B []int, expect int) {
	res := solve(x, y, A, B)
	if len(res) != expect {
		t.Errorf("Sample expect %d, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	x, y := 1, 4
	A := []int{1, 1, 1, 1, 2, 2, 2, 2, 3, 3, 3, 3, 4, 4, 4, 4, 4, 4, 4, 5}
	B := []int{1, 1, 1, 1, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 3, 3, 3, 3, 3, 3, 4, 4, 4, 4, 4, 4, 4, 4, 5, 5}
	expect := 20
	runSample(t, x, y, A, B, expect)
}
