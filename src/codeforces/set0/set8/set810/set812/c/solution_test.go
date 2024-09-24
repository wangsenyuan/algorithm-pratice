package main

import "testing"

func runSample(t *testing.T, s int, a []int, expect []int) {
	cnt, res := solve(a, s)

	if expect[0] != cnt || expect[1] != res {
		t.Fatalf("Sample expect %v, but got %d, %d", expect, cnt, res)
	}
}

func TestSample1(t *testing.T) {
	s := 11
	a := []int{2, 3, 5}
	expect := []int{2, 11}
	runSample(t, s, a, expect)
}

func TestSample2(t *testing.T) {
	s := 100
	a := []int{1, 2, 5, 6}
	expect := []int{4, 54}
	runSample(t, s, a, expect)
}

func TestSample3(t *testing.T) {
	s := 7
	a := []int{7}
	expect := []int{0, 0}
	runSample(t, s, a, expect)
}
