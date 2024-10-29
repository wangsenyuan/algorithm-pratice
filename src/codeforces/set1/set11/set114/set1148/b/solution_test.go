package main

import "testing"

func runSample(t *testing.T, ta int, tb int, k int, a []int, b []int, expect int) {
	res := solve(ta, tb, k, a, b)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	ta, tb, k := 1, 1, 2
	a := []int{1, 3, 5, 7}
	b := []int{1, 2, 3, 9, 10}
	expect := 11
	runSample(t, ta, tb, k, a, b, expect)
}

func TestSample2(t *testing.T) {
	ta, tb, k := 4, 4, 2
	a := []int{1, 10}
	b := []int{10, 20}
	expect := -1
	runSample(t, ta, tb, k, a, b, expect)
}

func TestSample3(t *testing.T) {
	ta, tb, k := 2, 3, 1
	a := []int{1, 999999998, 999999999, 1000000000}
	b := []int{3, 4, 1000000000}
	expect := 1000000003
	runSample(t, ta, tb, k, a, b, expect)
}
