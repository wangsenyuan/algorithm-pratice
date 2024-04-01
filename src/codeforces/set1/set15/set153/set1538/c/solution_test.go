package main

import "testing"

func runSample(t *testing.T, a []int, L int, R int, expect int) {
	res := solve(a, L, R)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{5, 1, 2}
	L, R := 4, 7
	expect := 2
	runSample(t, a, L, R, expect)
}

func TestSample2(t *testing.T) {
	a := []int{5, 1, 2, 4, 3}
	L, R := 5, 8
	expect := 7
	runSample(t, a, L, R, expect)
}
