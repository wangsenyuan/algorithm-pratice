package main

import (
	"sort"
	"testing"
)

func runSample(t *testing.T, n int) {
	res := solve(int64(n))

	a, b, c, d := res[0], res[1], res[2], res[3]

	sort.Slice(res, func(i, j int) bool {
		return res[i] < res[j]
	})
	if res[0] == 0 {
		t.Fatalf("Sample result %v, needs to > 0 for %d", res, n)
	}
	for i := 1; i < 3; i++ {
		if res[i] == res[i-1] {
			t.Fatalf("Sample result %v, needs to be distinct for %d", res, n)
		}
	}

	m := ((a & b) | c) ^ d

	if m != int64(n) {
		t.Fatalf("Sample resutl %v, got %d, but expect %d", []int64{a, b, c, d}, m, n)
	}
}

func TestSample1(t *testing.T) {
	for i := 0; i < 1000; i++ {
		runSample(t, i)
	}
}

func TestSample2(t *testing.T) {
	for i := 0; i < 31; i++ {
		runSample(t, 1<<i)
	}
}

func TestSample3(t *testing.T) {
	runSample(t, 3221225472)
}
