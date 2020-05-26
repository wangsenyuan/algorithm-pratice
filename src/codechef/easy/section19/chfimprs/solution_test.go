package main

import "testing"

func runSample(t *testing.T, n, m int, G [][]int, expect bool) {
	cnt := make(map[int]int)

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			cnt[G[i][j]]++
		}
	}

	res := solve(n, m, G)

	if res != expect {
		t.Fatalf("Sample expect %t, but got %t", expect, res)
	}

	if !res {
		return
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			cnt[G[i][j]]--
			if cnt[G[i][j]] < 0 {
				t.Fatalf("Sample expect %t, but got %t", expect, res)
			}
		}
	}
}

func TestSample1(t *testing.T) {
	n, m := 1, 2
	G := [][]int{{1, 1}}
	expect := true
	runSample(t, n, m, G, expect)
}

func TestSample2(t *testing.T) {
	n, m := 2, 2
	G := [][]int{{1, 1}, {2, 1}}
	expect := false
	runSample(t, n, m, G, expect)
}

func TestSample3(t *testing.T) {
	n, m := 3, 3
	G := [][]int{
		{1, 2, 3},
		{6, 5, 4},
		{7, 8, 9},
	}
	expect := false
	runSample(t, n, m, G, expect)
}
