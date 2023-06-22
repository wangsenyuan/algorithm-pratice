package main

import "testing"

func runSample(t *testing.T, persons [][]int, expect int) {
	res := int(solve(persons))
	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	persons := [][]int{
		{0, 0, 0, 1},
		{1, 0, 2, 0},
	}
	expect := 0
	runSample(t, persons, expect)
}

func TestSample2(t *testing.T) {
	persons := [][]int{
		{0, 0, 1, 1},
		{1, 1, 0, 0},
		{1, 0, 2, 0},
	}
	expect := 1
	runSample(t, persons, expect)
}

func TestSample3(t *testing.T) {
	persons := [][]int{
		{0, 0, 0, 1},
		{1, 0, 1, 2},
		{2, 0, 2, 3},
		{3, 0, 3, -5},
		{4, 0, 4, -5},
		{5, 0, 5, -5},
	}
	expect := 9
	runSample(t, persons, expect)
}
