package main

import "testing"

func runSample(t *testing.T, a, b []int, m int, expect int) {
	res := solve(a, b, m)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{7, 3, 6, 9}
	b := []int{4, 3, 8, 5}
	m := 2
	expect := 14
	runSample(t, a, b, m, expect)
}

func TestSample2(t *testing.T) {
	a := []int{6, 9, 7, 1, 8, 3}
	b := []int{5, 8, 8, 1, 4, 1}
	m := 2
	expect := 22
	runSample(t, a, b, m, expect)
}
