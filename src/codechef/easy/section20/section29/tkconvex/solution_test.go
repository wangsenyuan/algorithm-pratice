package main

import (
	"testing"
)

func runSample(t *testing.T, n int, k int, A []int, expect bool) {
	res := solve(n, k, A)

	var arr []int

	if len(res) > 0 {
		arr = make([]int, 2*k)
		for i := 0; i < 2*k; i++ {
			arr[i] = A[res[i]]
		}
	}

	if expect != (check(arr) > 0) {
		t.Errorf("Sample result %v, not correct, expect %t", res, expect)
	}
}

func TestSample1(t *testing.T) {
	n, k := 6, 3
	A := []int{1, 1, 1, 2, 2, 2}
	expect := true
	runSample(t, n, k, A, expect)
}

func TestSample2(t *testing.T) {
	n, k := 6, 3
	A := []int{1, 2, 3, 100, 200, 300}
	expect := false
	runSample(t, n, k, A, expect)
}
