package main

import (
	"testing"
)

func runSample(t *testing.T, a []int, k int) {
	n := len(a)
	sum := make([]int, n+1)
	var expect int
	var cnt int
	for i := 0; i < n; i++ {
		sum[i+1] = sum[i] + a[i]
		if a[i] == 0 {
			cnt++
			if cnt == k {
				expect = i + 1
			}
		}
	}

	ask := func(l int, r int) int {
		return sum[r] - sum[l-1]
	}

	res := solve(n, k, ask)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{1, 0, 1, 1, 0, 1}
	k := 2
	runSample(t, a, k)
}

func TestSample2(t *testing.T) {
	a := []int{1, 0, 1, 1, 0, 1}
	k := 1
	runSample(t, a, k)
}
