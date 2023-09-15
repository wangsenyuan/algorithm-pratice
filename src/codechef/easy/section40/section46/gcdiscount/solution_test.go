package main

import "testing"

func runSample(t *testing.T, A []int, B []int, expect int) {
	res := solve(A, B)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{24, 51, 96}
	B := []int{9, 44, 68}
	expect := 4
	runSample(t, A, B, expect)
}
