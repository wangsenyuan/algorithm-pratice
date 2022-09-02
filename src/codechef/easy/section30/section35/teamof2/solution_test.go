package main

import "testing"

func runSample(t *testing.T, n int, students [][]int, expect bool) {
	res := solve(n, students)

	if res != expect {
		t.Errorf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	students := [][]int{
		{3},
		{1},
		{1, 3, 4, 5},
	}
	expect := false
	runSample(t, n, students, expect)
}
