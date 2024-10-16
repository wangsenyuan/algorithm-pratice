package main

import (
	"slices"
	"testing"
)

func runSample(t *testing.T, segs [][]int, k int, expect []int) {
	res := solve(segs, k)

	if len(res) != len(expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}

	type pair struct {
		first  int
		second int
	}

	rm := make([]bool, len(segs)+1)
	for _, i := range res {
		rm[i] = true
	}

	var events []pair

	for i, cur := range segs {
		if rm[i+1] {
			continue
		}
		l, r := cur[0], cur[1]+1
		events = append(events, pair{l, i})
		events = append(events, pair{r, -i})
	}

	slices.SortFunc(events, func(a, b pair) int {
		if a.first == b.first {
			// 先处理离开的
			return a.second - b.second
		}
		return a.first - b.first
	})

	var active int

	for _, cur := range events {
		if cur.second > 0 {
			active++
		} else {
			active--
		}
		if active > k {
			t.Fatalf("Sample result %v, not correct", res)
		}
	}
}

func TestSample1(t *testing.T) {
	k := 2
	segs := [][]int{
		{11, 11},
		{9, 11},
		{7, 8},
		{8, 9},
		{7, 8},
		{9, 11},
		{7, 9},
	}
	expect := []int{1, 4, 5}
	runSample(t, segs, k, expect)
}

func TestSample2(t *testing.T) {
	k := 1
	segs := [][]int{
		{2, 3},
		{3, 3},
		{2, 3},
		{2, 2},
		{2, 3},
		{2, 3},
	}
	expect := []int{1, 3, 5, 6}
	runSample(t, segs, k, expect)
}
