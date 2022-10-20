package main

import "testing"

func runSample(t *testing.T, n int, R []int, B []int, expect int) {
	res := solve(n, R, B)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 1
	R := []int{101}
	B := []int{10}
	expect := 0
	runSample(t, n, R, B, expect)
}

func TestSample2(t *testing.T) {
	n := 2
	R := []int{1, 200}
	B := []int{199, 1}
	expect := 199
	runSample(t, n, R, B, expect)
}

func TestSample3(t *testing.T) {
	n := 4
	R := []int{1, 2, 3, 10}
	B := []int{3, 1, 3, 5}
	expect := 7
	runSample(t, n, R, B, expect)
}
