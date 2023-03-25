package main

import (
	"testing"
)

func runSample(t *testing.T, P []int, expect bool) {
	res := solve(P)

	if len(res) > 0 != expect {
		t.Fatalf("Sample expect %t, but got %v", expect, res)
	}

	if !expect {
		return
	}

	n := len(P)
	R := make([]int, n)
	R[n-1] = res[n-1]
	for i := n - 2; i >= 0; i-- {
		R[i] = min(res[i], R[i+1])
	}
	cur := res[0]

	for i := 1; i+1 < n; i++ {
		if cur > res[i] && res[i] > R[i+1] {
			t.Fatalf("Sample result %v, not correct", res)
		}
		cur = max(cur, res[i])
	}
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func TestSample1(t *testing.T) {
	P := []int{1, 2, 3}
	runSample(t, P, true)
}

func TestSample2(t *testing.T) {
	P := []int{1, 3, 2, 6, 5, 4}
	runSample(t, P, false)
}

func TestSample3(t *testing.T) {
	P := []int{4, 1, 3, 2}
	runSample(t, P, true)
}

func TestSample4(t *testing.T) {
	P := []int{3, 2, 1, 6, 7, 8, 5, 4}
	runSample(t, P, true)
}
