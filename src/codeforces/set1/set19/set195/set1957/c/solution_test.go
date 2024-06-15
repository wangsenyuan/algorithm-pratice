package main

import "testing"

func runSample(t *testing.T, n int, already [][]int, expect int) {
	res := solve(n, already)
	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 4
	already := [][]int{
		{1, 2},
	}
	expect := 3
	runSample(t, n, already, expect)
}

func TestSample2(t *testing.T) {
	n := 8
	already := [][]int{
		{7, 6},
	}
	expect := 331
	runSample(t, n, already, expect)
}
