package main

import "testing"

func runSample(t *testing.T, a [][]int, expect int) {
	cnt, res := solve(a)

	if cnt != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, cnt)
	}

	if expect < 0 {
		return
	}

	for _, cur := range res {
		i, j, k := cur[0]-1, cur[1]-1, cur[2]-1
		a[i][k], a[j][k] = a[j][k], a[i][k]
	}

	var x int
	for i := 0; i < len(a[0]); i++ {
		x += a[0][i]
	}

	for i := 1; i < len(a); i++ {
		var tmp int
		for j := 0; j < len(a[0]); j++ {
			tmp += a[i][j]
		}
		if tmp != x {
			t.Fatalf("Sample result %v, not correct", res)
		}
	}
}

func TestSample1(t *testing.T) {
	a := [][]int{
		{1, 1, 1, 0},
		{0, 0, 1, 0},
		{1, 0, 0, 1},
	}
	expect := 1
	runSample(t, a, expect)
}

func TestSample2(t *testing.T) {
	a := [][]int{
		{1, 0, 0},
		{0, 1, 1},
		{0, 0, 1},
		{0, 0, 0},
	}
	expect := 1
	runSample(t, a, expect)
}
