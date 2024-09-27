package main

import "testing"

func runSample(t *testing.T, n int, roads [][]int, pos []int, expect int) {
	res := solve(n, roads, pos)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}
func TestSample1(t *testing.T) {
	n := 5
	roads := [][]int{
		{1, 5, 1},
		{2, 4, 1},
	}
	pos := []int{3}
	expect := 5
	runSample(t, n, roads, pos, expect)
}

func TestSample2(t *testing.T) {
	n := 5
	roads := [][]int{
		{1, 5, 1},
		{2, 4, 1},
	}
	pos := []int{3, 4, 5}
	expect := 17
	runSample(t, n, roads, pos, expect)
}

func TestSample3(t *testing.T) {
	n := 3
	roads := [][]int{
		{1, 3, 1},
		{1, 3, 1},
	}
	pos := []int{2}
	expect := 4
	runSample(t, n, roads, pos, expect)
}

func TestSample4(t *testing.T) {
	n := 3
	roads := [][]int{
		{1, 3, 1},
		{1, 3, 1},
	}
	pos := []int{3}
	expect := 6
	runSample(t, n, roads, pos, expect)
}

func TestSample5(t *testing.T) {
	n := 10
	roads := [][]int{
		{6, 10, 74},
		{7, 9, 35},
		{3, 6, 63},
		{2, 4, 80},
		{2, 10, 78},
		{10, 10, 13},
		{4, 10, 16},
		{1, 2, 13},
		{3, 7, 17},
		{4, 6, 67},
	}
	pos := []int{9, 8, 10}
	expect := 635
	runSample(t, n, roads, pos, expect)
}
