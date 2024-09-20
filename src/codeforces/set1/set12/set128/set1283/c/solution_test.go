package main

import (
	"sort"
	"testing"
)

func runSample(t *testing.T, f []int) {
	res := solve(f)
	n := len(res)
	for i := 0; i < n; i++ {
		if f[i] > 0 && res[i] != f[i] {
			t.Fatalf("Sample result %v, should not change the original setting", res)
		}
		if res[i] == 0 {
			t.Fatalf("Sample result %v, having 0 value", res)
		}
		res[i]--
		if res[i] == i {
			t.Fatalf("Sample result %v, having self-present", res)
		}
	}
	vis := make([]bool, n)

	for i := 0; i < n; i++ {
		if !vis[i] {
			j := i
			for !vis[j] {
				vis[j] = true
				j = res[j]
			}
		}
	}

	for i := 0; i < n; i++ {
		if !vis[i] {
			t.Fatalf("Sample result %v, having %d no gift", res, i)
		}
	}

	sort.Ints(res)

	for i := 0; i < n; i++ {
		if res[i] != i {
			t.Fatalf("Sample result %v not correct", res)
		}
	}
}

func TestSample1(t *testing.T) {
	f := []int{5, 0, 0, 2, 4}
	runSample(t, f)
}

func TestSample2(t *testing.T) {
	f := []int{7, 0, 0, 1, 4, 0, 6}
	runSample(t, f)
}
