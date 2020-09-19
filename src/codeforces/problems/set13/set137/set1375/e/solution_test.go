package main

import (
	"sort"
	"testing"
)

func runSample(t *testing.T, n int, A []int) {
	B := make([]int, n)
	copy(B, A)

	invs := make(map[[2]int]bool)

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if A[i] > A[j] {
				invs[[...]int{i, j}] = true
			}
		}
	}

	res := solve(n, B)

	for _, x := range res {
		u, v := x[0], x[1]
		u--
		v--
		k := [...]int{u, v}
		if !invs[k] {
			t.Errorf("inversion %v not valid", x)
			return
		}
		delete(invs, k)
		A[u], A[v] = A[v], A[u]
	}

	if len(invs) > 0 {
		t.Errorf("Some inversions %v not in the list", invs)
		return
	}

	if !sort.IntsAreSorted(A) {
		t.Errorf("after all oprations, %v is not sorted", A)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	A := []int{3, 1, 2}
	runSample(t, n, A)
}

func TestSample2(t *testing.T) {
	n := 4
	A := []int{1, 8, 1, 6}
	runSample(t, n, A)
}

func TestSample3(t *testing.T) {
	n := 5
	A := []int{1, 1, 1, 2, 2}
	runSample(t, n, A)
}

func TestSample4(t *testing.T) {
	n := 6
	A := []int{2, 2, 1, 2, 2, 1, 2}
	runSample(t, n, A)
}
