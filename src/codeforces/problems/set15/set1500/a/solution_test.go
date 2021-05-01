package main

import "testing"

func runSample(t *testing.T, A []int, expect bool) {
	res := solve(A)
	if expect != (len(res) > 0) {
		t.Fatalf("Sample expect %t, but got %v", expect, res)
	}
	if expect {
		a, b, c, d := res[0], res[1], res[2], res[3]
		if A[a-1]+A[b-1] != A[c-1]+A[d-1] {
			t.Errorf("Sample result %v, not correct", res)
		}
	}
}

func TestSample1(t *testing.T) {
	A := []int{2, 1, 5, 2, 7, 4}
	runSample(t, A, true)
}

func TestSample2(t *testing.T) {
	A := []int{1, 3, 1, 9, 20}
	runSample(t, A, false)
}
