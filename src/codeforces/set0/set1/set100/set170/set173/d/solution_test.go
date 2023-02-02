package main

import "testing"

func runSample(t *testing.T, n int, E [][]int, expect bool) {
	res := solve(n, E)
	if len(res) > 0 != expect {
		t.Errorf("Sample expect %t, but got %v", expect, res)
		return
	}
	if !expect {
		return
	}
	if !check(res, n, E) {
		t.Errorf("Sample result %v, not correct", res)
	}
}

func check(res []int, n int, E [][]int) bool {
	k := n / 3

	cnt := make([]int, k+1)

	for _, e := range E {
		u, v := e[0], e[1]
		u--
		v--
		if res[u] == 0 || res[u] > k {
			return false
		}
		if res[v] == 0 || res[v] > k {
			return false
		}
		if res[u] == res[v] {
			return false
		}
	}

	for i := 0; i < n; i++ {
		cnt[res[i]]++
	}
	for i := 1; i <= k; i++ {
		if cnt[i] != 3 {
			return false
		}
	}
	return true
}

func TestSample1(t *testing.T) {
	n := 6
	E := [][]int{
		{1, 2},
		{1, 3},
		{4, 5},
		{4, 6},
	}
	expect := true
	runSample(t, n, E, expect)
}
