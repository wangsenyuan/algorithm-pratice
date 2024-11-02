package main

import "testing"

func runSample(t *testing.T, a []int, m int, expect int) {
	res := solve(a, m)

	if res != expect {
		t.Fatalf("Sample expect %d, but but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	m := 3
	a := []int{0, 0, 0, 1, 2}
	expect := 0
	runSample(t, a, m, expect)
}

func TestSample2(t *testing.T) {
	m := 7
	a := []int{0, 6, 1, 3, 2}
	expect := 1
	runSample(t, a, m, expect)
}

func TestSample3(t *testing.T) {
	m := 10
	a := []int{5, 0, 5, 9, 4, 6, 4, 5, 0, 0}
	expect := 6
	runSample(t, a, m, expect)
}
