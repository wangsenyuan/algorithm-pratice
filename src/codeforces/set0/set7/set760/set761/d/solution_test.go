package main

import "testing"

func runSample(t *testing.T, l int, r int, a []int, p []int, expect bool) {
	res := solve(l, r, a, p)

	if len(res) > 0 != expect {
		t.Fatalf("Sample expect %t, but got %v", expect, res)
	}

	if !expect {
		return
	}
	n := len(a)
	c := make([]int, n)
	for i := 0; i < n; i++ {
		c[i] = res[i] - a[i]
		if res[i] < l || res[i] > r {
			t.Fatalf("Sample result %v, exceeds the limit (%d %d)", res, l, r)
		}
	}
	pos := make([]int, n)
	for i, x := range p {
		pos[x-1] = i
	}

	for i := 1; i < n; i++ {
		if c[pos[i]] <= c[pos[i-1]] {
			t.Fatalf("Sample result %v,getting wrong result %v", res, c)
		}
	}
}

func TestSample1(t *testing.T) {
	l, r := 1, 5
	a := []int{1, 1, 1, 1, 1}
	p := []int{3, 1, 5, 4, 2}
	expect := true
	runSample(t, l, r, a, p, expect)
}

func TestSample2(t *testing.T) {
	l, r := 2, 9
	a := []int{3, 4, 8, 9}
	p := []int{3, 2, 1, 4}
	expect := true
	runSample(t, l, r, a, p, expect)
}

func TestSample3(t *testing.T) {
	l, r := 1, 5
	a := []int{1, 1, 1, 1, 1, 1}
	p := []int{2, 3, 5, 4, 1, 6}
	expect := false
	runSample(t, l, r, a, p, expect)
}
