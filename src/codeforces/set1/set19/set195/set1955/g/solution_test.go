package main

import "testing"

func runSample(t *testing.T, a [][]int, expect int) {
	res := solve(a)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := [][]int{
		{30, 20, 30},
		{15, 25, 40},
	}
	expect := 10
	runSample(t, a, expect)
}
