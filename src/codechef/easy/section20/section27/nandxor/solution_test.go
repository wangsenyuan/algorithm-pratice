package main

import "testing"

func runSample(t *testing.T, A []int, expect bool) {
	res := solve(A)

	if len(res) > 0 != expect {
		t.Fatalf("Sample expect %t, but got %v", expect, res)
	}

	if expect {
		x := A[res[0]-1] ^ A[res[1]-1]
		y := A[res[2]-1] ^ A[res[3]-1]
		if popcount(x) != popcount(y) {
			t.Errorf("Sample result %v, not correct", res)
		}
	}
}

func TestSample1(t *testing.T) {
	A := []int{1, 2, 5, 6}
	expect := true
	runSample(t, A, expect)
}
