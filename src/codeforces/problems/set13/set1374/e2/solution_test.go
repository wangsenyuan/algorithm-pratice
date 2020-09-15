package main

import (
	"testing"
)

func runSample(t *testing.T, n, m, k int, books [][]int, expect int64) {
	tot, res := solve(n, m, k, books)

	if tot != expect {
		t.Errorf("Sample expect %d, but got %d", expect, tot)
		return
	}
	if expect > 0 {
		mem := make(map[int]bool)
		for _, x := range res {
			mem[x] = true
		}
		if len(mem) != m {
			t.Errorf("Sample expect %d books but got %v", m, res)
			return
		}
		var k1, k2 int

		for i := 0; i < m; i++ {
			x := res[i] - 1
			if books[x][1] == 1 {
				k1++
			}
			if books[x][2] == 1 {
				k2++
			}
		}

		if k1 < k || k2 < k {
			t.Errorf("Sample result %v, not correct", res)
		}
	}
}

func TestSample1(t *testing.T) {
	n, m, k := 6, 3, 1
	books := [][]int{
		{6, 0, 0},
		{11, 1, 0},
		{9, 0, 1},
		{21, 1, 1},
		{10, 1, 0},
		{8, 0, 1},
	}
	var expect int64 = 24
	runSample(t, n, m, k, books, expect)
}

func TestSample2(t *testing.T) {
	n, m, k := 6, 3, 2
	books := [][]int{
		{6, 0, 0},
		{11, 1, 0},
		{9, 0, 1},
		{21, 1, 1},
		{10, 1, 0},
		{8, 0, 1},
	}
	var expect int64 = 39
	runSample(t, n, m, k, books, expect)
}

func TestSample3(t *testing.T) {
	n, m, k := 6, 6, 2
	books := [][]int{
		{6, 0, 0},
		{11, 1, 0},
		{9, 0, 1},
		{21, 1, 1},
		{10, 1, 0},
		{8, 0, 1},
	}
	var expect int64 = 65
	runSample(t, n, m, k, books, expect)
}

func TestSample4(t *testing.T) {
	n, m, k := 27, 5, 1
	books := [][]int{
		{232, 0, 1},
		{72, 0, 1},
		{235, 0, 1},
		{2, 0, 1},
		{158, 0, 0},
		{267, 0, 0},
		{242, 0, 1},
		{1, 0, 0},
		{64, 0, 0},
		{139, 1, 1},
		{250, 0, 1},
		{208, 0, 1},
		{127, 0, 1},
		{29, 0, 1},
		{53, 0, 1},
		{100, 0, 1},
		{52, 0, 1},
		{229, 0, 0},
		{1, 0, 1},
		{29, 0, 0},
		{17, 0, 1},
		{74, 0, 1},
		{211, 0, 1},
		{248, 0, 1},
		{15, 0, 0},
		{252, 0, 0},
		{159, 0, 1},
	}
	var expect int64 = 158
	runSample(t, n, m, k, books, expect)
}
