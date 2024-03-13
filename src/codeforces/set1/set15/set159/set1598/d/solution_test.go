package main

import "testing"

func runSample(t *testing.T, sessions [][]int, expect int) {
	res := solve(sessions)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	sessions := [][]int{
		{2, 4},
		{3, 4},
		{2, 1},
		{1, 3},
	}
	expect := 3
	runSample(t, sessions, expect)
}

func TestSample2(t *testing.T) {
	sessions := [][]int{
		{1, 5},
		{2, 4},
		{3, 3},
		{4, 2},
		{5, 1},
	}
	expect := 10
	runSample(t, sessions, expect)
}
