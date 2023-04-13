package main

import (
	"testing"
)

func runSample(t *testing.T, A [][]int, k int, expect int) {
	res := solve(A, k)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := [][]int{
		{2, 1, 0, 1, 3, 0, 1},
	}
	k := 5
	expect := 3
	runSample(t, A, k, expect)
}

func TestSample2(t *testing.T) {
	A := [][]int{
		{1, 0, 4, 0},
		{0, 2, 0, 3},
	}
	k := 3
	expect := 0
	runSample(t, A, k, expect)
}

func TestSample3(t *testing.T) {
	A := [][]int{
		{1, 0},
		{4, 1},
	}
	k := 7
	expect := -1
	runSample(t, A, k, expect)
}

func TestSample4(t *testing.T) {
	A := [][]int{
		{0, 2},
		{1, 0},
		{1, 0},
	}
	k := 3
	expect := 1
	runSample(t, A, k, expect)
}
