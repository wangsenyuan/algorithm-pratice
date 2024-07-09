package main

import "testing"

func runSample(t *testing.T, x []int, y []int, persons [][]int, expect int) {
	res := solve(x, y, persons)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	x := []int{0, 1000000}
	y := []int{0, 1000000}
	persons := [][]int{
		{1, 0},
		{1000000, 1},
		{999999, 1000000},
		{0, 999999},
	}
	expect := 2
	runSample(t, x, y, persons, expect)
}
func TestSample2(t *testing.T) {
	x := []int{0, 1, 2, 6, 1000000}
	y := []int{0, 4, 8, 1000000}
	persons := [][]int{
		{4, 4},
		{2, 5},
		{2, 2},
		{6, 3},
		{1000000, 1},
		{3, 8},
		{5, 8},
		{8, 8},
		{6, 8},
	}
	expect := 5
	runSample(t, x, y, persons, expect)
}
