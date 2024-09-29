package main

import "testing"

func runSample(t *testing.T, d []int, b []int, expect bool) {
	res := solve(d, b)

	if res != expect {
		t.Fatalf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	b := []int{5, 4, 3, 2, 1}
	d := []int{1, 1, 1, 1, 1}
	expect := true
	runSample(t, d, b, expect)
}

func TestSample2(t *testing.T) {
	b := []int{4, 3, 5, 1, 2, 7, 6}
	d := []int{4, 6, 6, 1, 6, 6, 1}
	expect := false
	runSample(t, d, b, expect)
}

func TestSample3(t *testing.T) {
	b := []int{4, 2, 5, 1, 3, 7, 6}
	d := []int{4, 6, 6, 1, 6, 6, 1}
	expect := true
	runSample(t, d, b, expect)
}
