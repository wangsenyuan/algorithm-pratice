package main

import (
	"testing"
)

func runSample(t *testing.T, sets [][]int, expect bool) {
	res := solve(sets)

	if len(res) > 0 != expect {
		t.Fatalf("Sample expect %t, but got %v", expect, res)
	}
	if !expect {
		return
	}
	a := res[0] - 1
	b := res[1] - 1
	cnt := make(map[int]int)
	for _, x := range sets[a] {
		cnt[x]++
	}
	var cnt2 int
	for _, x := range sets[b] {
		cnt2 += cnt[x]
	}
	if cnt2 < 2 {
		t.Fatalf("Sample result %v, not correct", res)
	}
}

func TestSample1(t *testing.T) {
	sets := [][]int{
		{1, 10},
		{1, 3, 5},
		{5, 4, 3, 2, 1},
		{10, 20, 30},
	}
	expect := true
	runSample(t, sets, expect)
}

func TestSample2(t *testing.T) {
	sets := [][]int{
		{1, 3, 5},
		{4, 3, 2},
	}
	expect := false
	runSample(t, sets, expect)
}
