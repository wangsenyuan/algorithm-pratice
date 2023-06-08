package main

import "testing"

func runSample(t *testing.T, n int, x int, expect bool) {
	res := solve(n, x)
	if len(res) > 0 != expect {
		t.Fatalf("Sample expect %t, but got %v", expect, res)
	}
	if !expect {
		return
	}

	set := make(map[int]bool)
	for i := 1; i <= n; i++ {
		set[i] = true
	}

	for _, cur := range res {
		a, b := cur[1], cur[2]
		if !set[a] || !set[b] {
			t.Fatalf("Sample result %d, have no corresponding nums", cur)
		}
		delete(set, a)
		delete(set, b)
		if cur[0] == 1 {
			set[a|b] = true
		} else {
			set[a^b] = true
		}
	}
	if len(set) != 1 || !set[x] {
		t.Fatalf("Sample get final result %v, not %d", set, x)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	x := 2
	expect := true
	runSample(t, n, x, expect)
}

func TestSample2(t *testing.T) {
	n := 2
	x := 2
	expect := false
	runSample(t, n, x, expect)
}

func TestSample3(t *testing.T) {
	n := 4
	x := 7
	expect := true
	runSample(t, n, x, expect)
}
