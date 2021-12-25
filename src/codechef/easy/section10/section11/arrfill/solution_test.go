package main

import "testing"

func runSample(t *testing.T, n, m int, P [][]int, expect int64) {
	res := solve(n, m, P)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, m := 10, 1
	P := [][]int{{5, 2}}
	var expect int64 = 25
	runSample(t, n, m, P, expect)
}

func TestSample2(t *testing.T) {
	n, m := 8, 2
	P := [][]int{{5, 2}, {6, 3}}
	var expect int64 = 41
	runSample(t, n, m, P, expect)
}

func TestSample3(t *testing.T) {
	n, m := 3, 2
	P := [][]int{{2, 2}, {1, 3}}
	var expect int64 = 5
	runSample(t, n, m, P, expect)
}
