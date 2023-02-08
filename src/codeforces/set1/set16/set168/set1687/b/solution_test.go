package main

import (
	"sort"
	"testing"
)

func runSample(t *testing.T, n int, E [][]int, expect int64) {
	P := make([]Pair, len(E))
	for i := 0; i < len(E); i++ {
		P[i] = Pair{int64(E[i][2]), i}
	}
	sort.Slice(P, func(i, j int) bool {
		return P[i].first < P[j].first
	})

	set := NewSet(n)

	ask2 := func(s string) int64 {
		return ask(s, E, P, set)
	}

	res := solve(n, len(E), ask2)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 5
	E := [][]int{
		{1, 2, 1},
		{2, 3, 2},
		{1, 3, 3},
		{4, 5, 4},
	}
	var expect int64 = 7
	runSample(t, n, E, expect)
}
