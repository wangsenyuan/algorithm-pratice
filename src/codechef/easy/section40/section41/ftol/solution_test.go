package main

import "testing"

func runSample(t *testing.T, n int, m int, special [][]int, expect int) {
	res := solve(n, m, special)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, m := 3, 3
	expect := 4
	runSample(t, n, m, nil, expect)
}

func TestSample2(t *testing.T) {
	n, m := 3, 3
	special := [][]int{
		{2, 2},
		{2, 3},
	}
	expect := 3
	runSample(t, n, m, special, expect)
}

func TestSample3(t *testing.T) {
	n, m := 1, 10
	special := [][]int{
		{1, 1},
	}
	expect := 9
	runSample(t, n, m, special, expect)
}
