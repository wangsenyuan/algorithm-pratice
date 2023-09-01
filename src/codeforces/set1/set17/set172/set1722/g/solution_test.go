package main

import (
	"sort"
	"testing"
)

func runSample(t *testing.T, n int) {
	res := solve(n)
	x := make([]int, 2)
	for i := 0; i < n; i++ {
		x[i&1] ^= res[i]
	}
	if x[0] != x[1] {
		t.Fatalf("Sample result %v, not correct", res)
	}

	sort.Ints(res)

	for i := 1; i < n; i++ {
		if res[i] == res[i-1] {
			t.Fatalf("Sample result %v, has duplicates", res)
		}
	}
}

func TestSample1(t *testing.T) {
	for i := 3; i <= 100; i++ {
		runSample(t, i)
	}
}

func TestSample2(t *testing.T) {
	runSample(t, 4)
}

func TestSample3(t *testing.T) {
	runSample(t, 6)
}
