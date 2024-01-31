package main

import "testing"

func runSample(t *testing.T, set []int, expect int) {
	s := make(map[int]bool)
	for _, num := range set {
		s[num] = true
	}
	ask := func(x int) int {
		if s[x] {
			panic("should add a non existing num")
		}
		s[x] = true
		x--
		for x >= 0 && !s[x] {
			x--
		}
		if s[x] {
			delete(s, x)
		}
		return x
	}

	mex := solve(set, ask)

	if mex != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, mex)
	}
}

func TestSample1(t *testing.T) {
	s := []int{1, 2, 3, 5, 7}
	expect := 4
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := []int{0, 1, 2}
	expect := 4
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := []int{5, 7, 57}
	expect := 1
	runSample(t, s, expect)
}
