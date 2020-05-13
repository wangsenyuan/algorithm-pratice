package main

import (
	"sort"
	"testing"
)

func runSample(t *testing.T, N, K int, A []int, expect bool) {
	ok, res := solve(N, K, A)

	if expect != ok {
		t.Fatalf("Sample %d %d %v, expect %t, but got %t", N, K, A, expect, ok)
	}

	if !expect {
		return
	}

	swap := func(i, j, k int) bool {
		if i >= j || j >= k || i >= k {
			return false
		}
		a, b, c := A[i], A[j], A[k]
		A[i] = c
		A[j] = a
		A[k] = b
		return true
	}

	for _, x := range res {
		if !swap(x[0]-1, x[1]-1, x[2]-1) {
			t.Fatalf("Sample Fail with %v", res)
		}
	}

	if !sort.IntsAreSorted(A) {
		t.Fatalf("Sample Fail with %v", res)
	}
}

func TestSample1(t *testing.T) {
	N, K := 4, 2
	A := []int{3, 2, 4, 1}
	expect := true
	runSample(t, N, K, A, expect)
}

func TestSample2(t *testing.T) {
	N, K := 4, 2
	A := []int{3, 4, 2, 1}
	expect := false
	runSample(t, N, K, A, expect)
}
