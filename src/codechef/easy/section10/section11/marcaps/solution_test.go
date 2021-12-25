package main

import "testing"

func runSample(t *testing.T, n int, A []int, expect bool) {
	X := make([]int, n)
	copy(X, A)
	res := solve(n, X)
	if res != expect {
		t.Errorf("Sample %d %v, expect %t, but got %t", n, A, expect, res)
	} else if expect {
		for i := 0; i < n; i++ {
			if X[i] == A[i] {
				t.Fatalf("Sample %d %v, expect %t, but got %t", n, A, expect, res)
			}
		}
		t.Logf("Sample %d %v, got %v", n, A, X)
	}
}

func TestSample1(t *testing.T) {
	n := 9
	A := []int{1, 1, 1, 2, 2, 2, 3, 3, 3}
	runSample(t, n, A, true)
}

func TestSample2(t *testing.T) {
	n := 2
	A := []int{1, 1}
	runSample(t, n, A, false)
}