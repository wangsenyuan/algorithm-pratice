package main

import "testing"

func runSample(t *testing.T, A []int, expect int) {
	res := solve(A)

	if res[2] != expect {
		t.Fatalf("Sample expect %d, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{7, 9}
	expect := 1
	runSample(t, A, expect)
}

func TestSample2(t *testing.T) {
	A := []int{5, 6, 7, 8, 9}
	expect := 5
	runSample(t, A, expect)
}
