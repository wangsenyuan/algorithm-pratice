package main

import "testing"

func runSample(t *testing.T, n int, E [][]int, expect string) {
	res := solve(n, E)

	if res != expect {
		t.Errorf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 2
	E := [][]int{{1, 2}}
	expect := "11"
	runSample(t, n, E, expect)
}

func TestSample2(t *testing.T) {
	n := 5
	E := [][]int{
		{1, 2},
		{1, 3},
		{3, 4},
		{3, 5},
	}
	expect := "00000"
	runSample(t, n, E, expect)
}
