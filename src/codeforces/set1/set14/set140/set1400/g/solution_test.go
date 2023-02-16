package main

import "testing"

func runSample(t *testing.T, n int, m int, L []int, R []int, E [][]int, expect int) {
	res := solve(n, m, L, R, E)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, m := 3, 0
	L := []int{1, 2, 1}
	R := []int{1, 3, 3}
	expect := 3
	runSample(t, n, m, L, R, nil, expect)
}

func TestSample2(t *testing.T) {
	n, m := 3, 1
	L := []int{1, 2, 1}
	R := []int{1, 3, 3}
	E := [][]int{{2, 3}}
	expect := 2
	runSample(t, n, m, L, R, E, expect)
}
