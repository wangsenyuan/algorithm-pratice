package main

import (
	"sort"
	"testing"
)

func runSample(t *testing.T, A []int, expect bool) {
	arr := make([]int, len(A))
	copy(arr, A)

	ok, res := solve(A)

	if ok != expect || len(res) > 4 {
		t.Fatalf("Sample expect %t, but got %t", expect, ok)
	}

	if !ok {
		return
	}
	n := len(A)
	for _, vs := range res {
		tmp := arr[vs[n-1]-1]
		for i := n - 1; i > 0; i-- {
			arr[vs[i]-1] = arr[vs[i-1]-1]
		}
		arr[vs[0]-1] = tmp
	}

	if !sort.IntsAreSorted(arr) {
		t.Fatalf("Sample result %v, not correct", res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{1, 2, 3}
	expect := true
	runSample(t, A, expect)
}

func TestSample2(t *testing.T) {
	A := []int{3, 4, 1, 2}
	expect := true
	runSample(t, A, expect)
}

func TestSample3(t *testing.T) {
	A := []int{4, 1, 5, 2, 3}
	expect := false
	runSample(t, A, expect)
}
