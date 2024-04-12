package main

import "testing"

func runSample(t *testing.T, a []int, expect int) {
	ask := func(i int) int {
		return a[i-1]
	}

	res := solve(len(a), ask)

	if expect < 0 && res > 0 {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
	if expect < 0 {
		return
	}
	if res < 0 {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
	res--
	if a[res] != a[len(a)/2+res] {
		t.Fatalf("Sample result %d, not correct", res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{1, 2, 1, 2, 3, 4, 3, 2}
	expect := 4
	runSample(t, a, expect)
}

func TestSample2(t *testing.T) {
	a := []int{1, 2, 3, 2, 1, 0}
	expect := -1
	runSample(t, a, expect)
}
