package main

import "testing"

func runSample(t *testing.T, n int, L []int, R []int, expect int) {
	res := int(solve(n, L, R))

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 1
	L := []int{3}
	R := []int{4}
	expect := 0
	runSample(t, n, L, R, expect)
}

func TestSample2(t *testing.T) {
	n := 4
	L := []int{2, 1, 3, 2}
	R := []int{4, 2, 4, 3}
	expect := 5
	runSample(t, n, L, R, expect)
}

func TestSample3(t *testing.T) {
	n := 7
	L := []int{3, 10, 5, 2, 8, 3, 9}
	R := []int{8, 6, 11, 5, 9, 13, 7}
	expect := 32
	runSample(t, n, L, R, expect)
}
