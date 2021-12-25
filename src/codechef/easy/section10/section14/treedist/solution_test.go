package main

import "testing"

func runSample(t *testing.T, n int, E [][]int, expect int) {
	res := solve(n, E)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 6
	E := [][]int{{1, 2},
		{1, 3},
		{3, 4},
		{2, 5},
		{2, 6},
	}
	expect := 54
	runSample(t, n, E, expect)

}
