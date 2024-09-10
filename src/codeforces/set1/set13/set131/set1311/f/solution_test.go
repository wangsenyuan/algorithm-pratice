package main

import "testing"

func runSample(t *testing.T, x []int, v []int, expect int) {
	res := solve(x, v)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	x := []int{1, 3, 2}
	v := []int{-100, 2, 3}
	expect := 3
	runSample(t, x, v, expect)
}

func TestSample2(t *testing.T) {
	x := []int{2, 1, 4, 3, 5}
	v := []int{2, 2, 2, 3, 4}
	expect := 19
	runSample(t, x, v, expect)
}

func TestSample3(t *testing.T) {
	x := []int{2, 1}
	v := []int{-3, 0}
	expect := 0
	runSample(t, x, v, expect)
}
