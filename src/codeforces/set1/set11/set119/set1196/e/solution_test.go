package main

import (
	"testing"
)

func runSample(t *testing.T, b int, w int, expect bool) {
	res := solve(b, w)

	if len(res) > 0 != expect {
		t.Fatalf("Sample expect %t, but got %v", expect, res)
	}
	if !expect {
		return
	}

	type pair struct {
		first  int
		second int
	}
	marked := make(map[pair]bool)
	cnt := make([]int, 2)
	for _, cur := range res {
		r, c := cur[0], cur[1]
		r--
		c--
		cnt[(r+c)&1]++
		k := pair{r, c}
		if marked[k] {
			t.Fatalf("Sample result %v, has overlapped cell at %d %d", res, r, c)
		}
		marked[k] = true
		// 怎么验证它们是连通的呢？
	}

	if cnt[0] != w || cnt[1] != b {
		t.Fatalf("Sample result %v, not correct", res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 1, 1, true)
}

func TestSample2(t *testing.T) {
	runSample(t, 1, 4, true)
}

func TestSample3(t *testing.T) {
	runSample(t, 2, 5, true)
}
