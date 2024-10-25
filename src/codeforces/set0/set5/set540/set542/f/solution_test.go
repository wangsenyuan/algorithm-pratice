package main

import "testing"

func runSample(t *testing.T, tot int, tasks [][]int, expect int) {
	res := solve(tot, tasks)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	tot := 5
	tasks := [][]int{
		{1, 1},
		{1, 1},
		{2, 2},
		{3, 3},
		{4, 4},
	}
	expect := 11
	runSample(t, tot, tasks, expect)
}

func TestSample2(t *testing.T) {
	tot := 5
	tasks := [][]int{
		{4, 1},
		{4, 2},
		{4, 3},
		{4, 4},
		{4, 5},
	}
	expect := 9
	runSample(t, tot, tasks, expect)
}

func TestSample3(t *testing.T) {
	tot := 2
	tasks := [][]int{
		{1, 1},
		{2, 10},
	}
	expect := 10
	runSample(t, tot, tasks, expect)
}
