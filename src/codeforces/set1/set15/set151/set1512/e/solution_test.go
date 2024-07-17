package main

import (
	"sort"
	"testing"
)

func runSample(t *testing.T, n int, l int, r int, s int, expect bool) {
	res := solve(n, l, r, s)
	if len(res) > 0 != expect {
		t.Fatalf("Sample expect %t, but got %v", expect, res)
	}

	if !expect {
		return
	}

	var sum int
	for i := l - 1; i < r; i++ {
		sum += res[i]
	}

	if sum != s {
		t.Fatalf("Sample result %v, not getting expect sum %d", res, s)
	}

	sort.Ints(res)

	for i := 0; i < n; i++ {
		if res[i] != i+1 {
			t.Fatalf("Sample result %v, not a permuation", res)
		}
	}
}

func TestSample1(t *testing.T) {
	n, l, r, s := 5, 2, 3, 5
	expect := true
	runSample(t, n, l, r, s, expect)
}

func TestSample2(t *testing.T) {
	n, l, r, s := 5, 3, 4, 1
	expect := false
	runSample(t, n, l, r, s, expect)
}

func TestSample3(t *testing.T) {
	n, l, r, s := 3, 1, 2, 4
	expect := true
	runSample(t, n, l, r, s, expect)
}

func TestSample4(t *testing.T) {
	n, l, r, s := 2, 2, 2, 2
	expect := true
	runSample(t, n, l, r, s, expect)
}

func TestSample5(t *testing.T) {
	n, l, r, s := 1, 1, 1, 1
	expect := true
	runSample(t, n, l, r, s, expect)
}

func TestSample6(t *testing.T) {
	n, l, r, s := 3, 1, 3, 6
	expect := true
	runSample(t, n, l, r, s, expect)
}

func TestSample7(t *testing.T) {
	n, l, r, s := 3, 1, 1, 2
	expect := true
	runSample(t, n, l, r, s, expect)
}
