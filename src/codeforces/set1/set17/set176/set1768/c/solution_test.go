package main

import (
	"sort"
	"testing"
)

func runSample(t *testing.T, a []int, expect bool) {
	ok, x, y := solve(a)

	if ok != expect {
		t.Fatalf("Sample expect %t but got %t", expect, ok)
	}
	if !expect {
		return
	}

	for i := 0; i < len(a); i++ {
		if a[i] != max(x[i], y[i]) {
			t.Fatalf("Sample result %v %v, not correct", x, y)
		}
	}
	n := len(x)
	sort.Ints(x)

	if x[0] != 1 || x[n-1] != n {
		t.Fatalf("Sample result %v %v, not correct", x, y)
	}

	sort.Ints(y)
	if y[0] != 1 || y[n-1] != n {
		t.Fatalf("Sample result %v %v, not correct", x, y)
	}
}

func TestSample1(t *testing.T) {
	a := []int{5, 3, 4, 2, 5}
	expect := true
	runSample(t, a, expect)
}

func TestSample2(t *testing.T) {
	a := []int{1, 1}
	expect := false
	runSample(t, a, expect)
}

func TestSample3(t *testing.T) {
	a := []int{3, 3, 3}
	expect := false
	runSample(t, a, expect)
}
