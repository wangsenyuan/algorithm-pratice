package main

import (
	"sort"
	"testing"
)

func runSample(t *testing.T, A []int, B []int, expect bool) {
	a := make([]int, len(A))
	b := make([]int, len(B))
	copy(a, A)
	copy(b, B)
	ok, res := solve(A, B)

	if ok != expect {
		t.Fatalf("Sample expect %t, but got %t", expect, ok)
	}
	if ok {
		for _, x := range res {
			i, j := x[0], x[1]
			i--
			j--
			a[i], a[j] = a[j], a[i]
			b[i], b[j] = b[j], b[i]
		}

		if !sort.IntsAreSorted(a) || !sort.IntsAreSorted(b) {
			t.Fatalf("result %v not sort A or B correctly", res)
		}
	}
}

func TestSample1(t *testing.T) {
	A := []int{1, 2}
	B := []int{1, 2}
	expect := true
	runSample(t, A, B, expect)
}

func TestSample2(t *testing.T) {
	A := []int{1, 2}
	B := []int{2, 1}
	expect := false
	runSample(t, A, B, expect)
}

func TestSample3(t *testing.T) {
	A := []int{2, 3, 1, 2}
	B := []int{2, 3, 2, 3}
	expect := true
	runSample(t, A, B, expect)
}
