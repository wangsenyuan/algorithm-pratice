package main

import "testing"

func runSample(t *testing.T, A []int, expect bool) {
	res := solve(A)

	if res != expect {
		t.Fatalf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{20, 16, 3, 13, 9, 19, 3, 17, 5, 12, 13, 6, 0, 4, 20, 5, 11, 17, 9, 9}
	expect := true
	runSample(t, A, expect)
}
