package main

import (
	"sort"
	"testing"
)

func runSample(t *testing.T, a []int, expect bool) {
	n := len(a)
	p := make([]int, n)
	copy(p, a)
	ok, res := solve(a)

	if ok != expect {
		t.Fatalf("Sample expect %t, but got %t, %v", expect, ok, res)
	}

	if !ok {
		return
	}

	if len(res) > n*5/2 {
		t.Fatalf("Sample result %v, took too much times to reverse", res)
	}

	for _, x := range res {
		if x%2 == 0 {
			t.Fatalf("Sample result %v, not correct, it has even length reverse ", res)
		}
		for i, j := 0, x-1; i < j; i, j = i+1, j-1 {
			p[i], p[j] = p[j], p[i]
		}
	}

	if !sort.IntsAreSorted(p) {
		t.Fatalf("Sample result %v, not correct", res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{1, 2, 3}
	expect := true
	runSample(t, a, expect)
}
func TestSample2(t *testing.T) {
	a := []int{3, 4, 5, 2, 1}
	expect := true
	runSample(t, a, expect)
}

func TestSample3(t *testing.T) {
	a := []int{2, 1, 3}
	expect := false
	runSample(t, a, expect)
}
