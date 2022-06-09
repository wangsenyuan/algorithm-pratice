package main

import "testing"

func runSample(t *testing.T, n int, C []int, V []int, expect int) {
	res := int(solve(n, C, V))

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 6
	C := []int{1, 2, 5, 1, 1, 3}
	V := []int{38, 46, 66, 64, 59, 69}
	expect := 69
	runSample(t, n, C, V, expect)
}

func TestSample2(t *testing.T) {
	n := 6
	C := []int{2, 5, 4, 1, 1, 2}
	V := []int{49, 24, 17, 54, 23, 50}
	expect := 31
	runSample(t, n, C, V, expect)
}

func TestSample3(t *testing.T) {
	n := 2
	C := []int{4, 2}
	V := []int{32, 78}
	expect := 0
	runSample(t, n, C, V, expect)
}
