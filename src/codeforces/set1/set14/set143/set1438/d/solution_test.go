package main

import (
	"sort"
	"testing"
)

func runSample(t *testing.T, A []int, expect bool) {
	cnt, res := solve(A)
	if cnt > 0 != expect {
		t.Fatalf("Sample expect %t, but got %d, %v", expect, cnt, res)
	}

	if !expect {
		return
	}
	arr := make([]int, len(A)+1)
	copy(arr[1:], A)
	for _, cur := range res {
		a, b, c := cur[0], cur[1], cur[2]
		x := arr[a] ^ arr[b] ^ arr[c]
		arr[a] = x
		arr[b] = x
		arr[c] = x
	}
	if !sort.IntsAreSorted(arr[1:]) || arr[1] != arr[len(A)] {
		t.Fatalf("Sample result %v not correct, get a array %v", res, arr[1:])
	}
}

func TestSample1(t *testing.T) {
	A := []int{4, 2, 1, 7, 2}
	expect := true
	runSample(t, A, expect)
}
