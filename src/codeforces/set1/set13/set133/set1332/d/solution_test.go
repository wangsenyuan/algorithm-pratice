package main

import "testing"

func runF(t *testing.T, a [][]int, expect int) {
	res := f(a)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestF1(t *testing.T) {
	a := [][]int{
		{7, 3, 3, 1},
		{4, 8, 3, 6},
		{7, 7, 7, 3},
	}
	expect := 2
	runF(t, a, expect)
}
