package main

import "testing"

func runSample(t *testing.T, A, B, C []int, k1, k2 int, expect int) {
	res := int(solve(A, B, C, k1, k2))

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{10, 4, 8, 5}
	B := []int{9, 7, 6, 6}
	C := []int{5, 5, 7, 11}
	k1, k2 := 2, 1
	expect := 36
	runSample(t, A, B, C, k1, k2, expect)
}

func TestSample2(t *testing.T) {
	A := []int{12, 44, 32, 12, 32}
	B := []int{43, 32, 12, 32, 31}
	C := []int{34, 12, 43, 23, 41}
	k1, k2 := 3, 4
	expect := 203
	runSample(t, A, B, C, k1, k2, expect)
}
