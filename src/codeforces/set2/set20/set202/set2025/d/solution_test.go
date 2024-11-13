package main

import "testing"

func runSample(t *testing.T, m int, r []int, expect int) {
	res := solve(m, r)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	m := 5
	r := []int{0, 1, 0, 2, 0, -3, 0, -4, 0, -5}
	expect := 3
	runSample(t, m, r, expect)
}

func TestSample2(t *testing.T) {
	m := 1
	r := []int{1, -1, 0}
	expect := 0
	runSample(t, m, r, expect)
}

func TestSample3(t *testing.T) {
	m := 3
	r := []int{0, 0, 1, 0, 2, -3, -2, -2, 1}
	expect := 4
	runSample(t, m, r, expect)
}
