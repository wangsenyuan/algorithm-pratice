package main

import "testing"

func runSample(t *testing.T, n int, k int, a []int, b []int, expect int) {
	res := solve(n, k, a, b)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{1, 2}
	b := []int{3, 4}
	n, k := 3, 5
	expect := 7
	runSample(t, n, k, a, b, expect)
}

func TestSample2(t *testing.T) {
	a := []int{1, 1}
	b := []int{999999999, 999999999}
	n, k := 2, 1000000000
	expect := 2000000000
	runSample(t, n, k, a, b, expect)
}

func TestSample3(t *testing.T) {
	a := []int{5, 10}
	b := []int{7, 8}
	n, k := 10, 3
	expect := 0
	runSample(t, n, k, a, b, expect)
}
