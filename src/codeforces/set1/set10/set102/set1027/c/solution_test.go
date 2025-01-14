package main

import (
	"sort"
	"testing"
)

func runSample(t *testing.T, a []int, expect []int) {
	res := solve(a)

	check := func(arr []int) pair {
		sort.Ints(arr)
		h := arr[0]
		w := arr[3]
		p := 2 * (h + w)
		p *= p
		s := h * w
		c := gcd(p, s)
		return pair{p / c, s / c}
	}

	x := check(expect)
	y := check(res)

	if x != y {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{7, 2, 2, 7}
	expect := []int{7, 2, 2, 7}
	runSample(t, a, expect)
}

func TestSample2(t *testing.T) {
	a := []int{2, 8, 1, 4, 8, 2, 1, 5}
	expect := []int{1, 2, 2, 1}
	runSample(t, a, expect)
}

func TestSample3(t *testing.T) {
	a := []int{5, 5, 5, 5, 5}
	expect := []int{5, 5, 5, 5}
	runSample(t, a, expect)
}

func TestSample4(t *testing.T) {
	a := []int{24, 8455, 24, 9508, 9508, 8455}
	expect := []int{8455, 8455, 9508, 9508}
	runSample(t, a, expect)
}
