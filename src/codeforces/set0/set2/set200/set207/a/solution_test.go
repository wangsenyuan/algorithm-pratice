package main

import "testing"

func runSample(t *testing.T, n int, scientists [][]int, expect int) {
	bad, res := solve(n, scientists)

	if bad != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, bad)
	}

	val := make([]int, n)
	pos := make([]int, n)

	for i := 0; i < n; i++ {
		val[i] = scientists[i][1]
	}
	var cnt int
	for i, cur := range res {
		if i > 0 && cur[0] < res[i-1][0] {
			cnt++
		}
		v, j := int(cur[0]), cur[1]-1
		if val[j] != v {
			t.Fatalf("Sample result %v, not correct", res)
		}
		pos[j]++
		scientist := scientists[j]
		k := scientist[0]
		if pos[j] == k {
			continue
		}
		// for next check
		v = (v*scientist[2] + scientist[3]) % scientist[4]
		val[j] = v
	}

	if bad != cnt {
		t.Fatalf("Sample result %v, not getting correct %d", res, expect)
	}
}

func TestSample1(t *testing.T) {
	n := 2
	scientists := [][]int{
		{2, 1, 1, 1, 10},
		{2, 3, 1, 1, 10},
	}
	expect := 0
	runSample(t, n, scientists, expect)
}

func TestSample2(t *testing.T) {
	n := 2
	scientists := [][]int{
		{3, 10, 2, 3, 1000},
		{3, 100, 1, 999, 1000},
	}
	expect := 2
	runSample(t, n, scientists, expect)
}

func TestSample3(t *testing.T) {
	n := 2
	scientists := [][]int{
		{5, 737247526, 385663770, 400009853, 926128972},
		{5, 134299091, 206672784, 206760990, 541043536},
	}
	expect := 3
	runSample(t, n, scientists, expect)
}
