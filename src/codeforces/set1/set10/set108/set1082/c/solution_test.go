package main

import "testing"

func runSample(t *testing.T, m int, students [][]int, expect int) {
	res := solve(m, students)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	m := 3
	students := [][]int{
		{2, 6},
		{3, 6},
		{2, 5},
		{3, 5},
		{1, 9},
		{3, 1},
	}
	expect := 22
	runSample(t, m, students, expect)
}

func TestSample2(t *testing.T) {
	m := 2
	students := [][]int{
		{1, -1},
		{1, -5},
		{2, -1},
		{2, -1},
		{1, -10},
	}
	expect := 0
	runSample(t, m, students, expect)
}
