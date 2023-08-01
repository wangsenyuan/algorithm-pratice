package main

import "testing"

func runSample(t *testing.T, n int, m int, days [][]int, expect bool) {
	res := solve(n, m, days)
	if len(res) > 0 != expect {
		t.Fatalf("Sample expect %t, but got %v", expect, res)
	}

	if !expect {
		return
	}
	cnt := make(map[int]int)
	for i := 0; i < m; i++ {
		x := res[i]
		ok := contains(days[i], x-1)
		if !ok {
			t.Fatalf("Sample result %v, ans for day %d not in the list %v", res, i, days[i])
		}
		cnt[x]++
	}

	for k, v := range cnt {
		if v > (m+1)/2 {
			t.Fatalf("Sample result %v, %d occures %d times > %d", res, k, v, (m+1)/2)
		}
	}
}

func contains(arr []int, x int) bool {
	for _, y := range arr {
		if x == y {
			return true
		}
	}
	return false
}

func TestSample1(t *testing.T) {
	n, m := 4, 6
	days := [][]int{
		{1},
		{1, 2},
		{1, 2, 3},
		{1, 2, 3, 4},
		{2, 3},
		{3},
	}
	expect := true
	runSample(t, n, m, days, expect)
}

func TestSample2(t *testing.T) {
	n, m := 2, 2
	days := [][]int{
		{1},
		{1},
	}
	expect := false
	runSample(t, n, m, days, expect)
}
