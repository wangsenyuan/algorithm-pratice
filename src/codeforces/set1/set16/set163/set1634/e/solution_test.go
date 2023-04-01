package main

import "testing"

func runSample(t *testing.T, arrs [][]int, expect bool) {
	ok, res := solve(arrs)

	if ok != expect {
		t.Fatalf("Sample expect %t, but got %t", expect, ok)
	}
	if !expect {
		return
	}
	cnt := make(map[int]int)
	for i := 0; i < len(arrs); i++ {
		var tmp int
		for j := 0; j < len(arrs[i]); j++ {
			if res[i][j] == 'L' {
				tmp++
				cnt[arrs[i][j]]++
			} else {
				cnt[arrs[i][j]]--
			}
		}
		if tmp*2 != len(arrs[i]) {
			t.Fatalf("Sample result %v, not correct", res)
		}
	}
	for _, v := range cnt {
		if v != 0 {
			t.Fatalf("Sample result %v, not correct", res)
		}
	}
}

func TestSample1(t *testing.T) {
	arrs := [][]int{
		{1, 2},
		{1, 2, 3, 3},
		{1, 1, 2, 2, 3, 3},
	}
	expect := true
	runSample(t, arrs, expect)
}
