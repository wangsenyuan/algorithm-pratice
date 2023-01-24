package main

import (
	"sort"
	"testing"
)

func runSample(t *testing.T, A []int, expect bool) {
	arr := make([]int, len(A))
	copy(arr, A)
	ok, res := solve(A)

	if expect != ok {
		t.Fatalf("Sample expect %t, but got %t", expect, ok)
	}

	if !expect {
		return
	}

	for _, swap := range res {
		i, j := swap[0], swap[1]
		i--
		j--
		if !checkLuck(arr[i]) && !checkLuck(arr[j]) {
			t.Fatalf("Sample result %v, swap non-lucky numbers", res)
		}
		arr[i], arr[j] = arr[j], arr[i]
	}
	if !sort.IntsAreSorted(arr) {
		t.Fatalf("Sample result %v, not getting sorted result: %v", res, arr)
	}
}

func TestSample1(t *testing.T) {
	A := []int{4, 2, 1}
	expect := true
	runSample(t, A, expect)
}

func TestSample2(t *testing.T) {
	A := []int{77, 66, 55, 44, 33, 22, 11}
	expect := true
	runSample(t, A, expect)
}
