package main

import "testing"

func runSample(t *testing.T, persons [][]int, expect int) {
	res := solve(persons)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	persons := [][]int{
		{9, 3, 3},
		{4, 3, 1},
		{2, 3, 3},
	}
	expect := 15
	runSample(t, persons, expect)
}

func TestSample2(t *testing.T) {
	persons := [][]int{
		{2, 4, 6},
		{3, 5, 5},
		{5, 6, 4},
	}
	expect := 24
	runSample(t, persons, expect)
}
