package main

import (
	"slices"
	"testing"
)

func runSample(t *testing.T, a []int, k int, m int) {
	n := len(a)
	tmp := make([]pair, k)
	var cnt int
	ask := func(arr []int) (int, int) {
		cnt++
		if cnt > n {
			t.Fatal("asked too much")
		}
		for i := 0; i < k; i++ {
			tmp[i] = pair{a[arr[i]-1], arr[i]}
		}
		slices.SortFunc(tmp, func(x, y pair) int {
			return x.first - y.first
		})
		return tmp[m-1].second, tmp[m-1].first
	}

	ans := solve(n, k, ask)

	if ans != m {
		t.Fatalf("Sample result got wrong %d, instead of %d", ans, m)
	}
}

func TestSample1(t *testing.T) {
	a := []int{2, 0, 1, 9}
	k := 3
	m := 3
	runSample(t, a, k, m)
}
