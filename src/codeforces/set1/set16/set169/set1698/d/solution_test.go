package main

import (
	"sort"
	"testing"
)

func runSample(t *testing.T, arr []int) {
	var cnt int
	ask := func(l int, r int) []int {
		tmp := make([]int, r-l+1)
		copy(tmp, arr[l-1:r])
		sort.Ints(tmp)
		cnt++
		return tmp
	}

	res := solve(len(arr), ask)

	if arr[res-1] != res {
		t.Fatalf("Sample got wrong result %d", res)
	}
	if cnt > 15 {
		t.Fatalf("Sample uses %d tries > 15", cnt)
	}
}

func TestSample1(t *testing.T) {
	arr := []int{4, 2, 5, 1, 3}
	runSample(t, arr)
}

func TestSample2(t *testing.T) {
	arr := []int{1, 3, 2}
	runSample(t, arr)
}
