package main

import (
	"sort"
	"testing"
)

func runSample(t *testing.T, a []int, expect int) {
	n, res := solve(a)
	if n != expect {
		t.Fatalf("Sample %v, expect %d, but got %d", a, expect, n)
	}
	if expect < 0 {
		return
	}
	arr := make([][]int, n)
	for i, j := range res {
		arr[j-1] = append(arr[j-1], a[i])
	}
	for _, cur := range arr {
		sort.Ints(cur)
		for i := 0; i < len(cur); i++ {
			if cur[i] != i+1 {
				t.Fatalf("Sample %v, expect %d, but got %d", a, i, cur[i])
			}
		}
	}
}

func TestSample1(t *testing.T) {
	a := []int{1, 2, 3, 1, 2, 1, 4, 2, 5}
	expect := 3
	runSample(t, a, expect)
}

func TestSample2(t *testing.T) {
	a := []int{1, 2, 2, 3}
	expect := -1
	runSample(t, a, expect)
}
