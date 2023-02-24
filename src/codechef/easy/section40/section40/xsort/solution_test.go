package main

import (
	"sort"
	"testing"
)

func runSample(t *testing.T, a []int) {
	n := len(a)
	A := make([]int, n)
	copy(A, a)
	res := solve(a)
	if len(res) > 20*n {
		t.Fatalf("Sample expect result to be less than < 20 * %d", n)
	}
	for _, x := range res {
		i, j := x[0], x[1]
		i--
		j--
		y := A[i] ^ A[j]
		A[i] = y
		A[j] = y
	}
	if !sort.IntsAreSorted(A) {
		t.Errorf("Sample result %v, not correct", res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{34, 40, 20, 30}
	runSample(t, a)
}
