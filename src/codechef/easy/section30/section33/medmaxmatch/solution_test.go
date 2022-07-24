package main

import "testing"

func runSample(t *testing.T, A []int, B []int, expect int) {
	res := solve(A, B)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{10, 4, 93, 5, 16}
	B := []int{4, 34, 62, 6, 26}
	expect := 50
	runSample(t, A, B, expect)
}
