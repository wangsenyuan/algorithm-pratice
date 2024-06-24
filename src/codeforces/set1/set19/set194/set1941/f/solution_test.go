package main

import "testing"

func runSample(t *testing.T, a []int, d []int, f []int, expect int) {
	res := solve(a, d, f)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{5, 10, 15, 20, 26}
	d := []int{11, 14, 16, 13, 8}
	f := []int{16, 4, 5, 3, 1}
	expect := 5
	runSample(t, a, d, f, expect)
}

func TestSample2(t *testing.T) {
	a := []int{1, 4, 7, 10, 18, 21, 22}
	d := []int{2, 3, 5, 7, 4, 2}
	f := []int{6, 8, 9, 3, 2}
	expect := 4
	runSample(t, a, d, f, expect)
}

func TestSample3(t *testing.T) {
	a := []int{4, 21}
	d := []int{4, 15, 14, 5}
	f := []int{20, 1, 15, 1, 12, 5, 11}
	expect := 11
	runSample(t, a, d, f, expect)
}
