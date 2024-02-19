package main

import "testing"

func runSample(t *testing.T, m int, a []int, b []int, expect int) {
	res := solve(m, a, b)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	m := 4
	a := []int{1}
	b := []int{3, 2}
	expect := 2
	runSample(t, m, a, b, expect)
}

func TestSample2(t *testing.T) {
	m := 7
	a := []int{5, 1, 5}
	b := []int{3, 8, 3, 3}
	expect := 12
	runSample(t, m, a, b, expect)
}

func TestSample3(t *testing.T) {
	m := 4
	a := []int{4, 3, 3, 2, 2, 1, 1}
	b := []int{1, 1, 1, 1, 3, 3, 3, 3}
	expect := 16
	runSample(t, m, a, b, expect)
}
