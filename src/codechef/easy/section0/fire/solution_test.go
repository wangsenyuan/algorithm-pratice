package main

import "testing"

func runSample(t *testing.T, n int, s int, neighbors [][]int, important []int, expect bool) {
	res := solve(n, s, neighbors, important)
	if res != expect {
		t.Errorf("this sample %d, %d, %v, %v should give answer %t, but got %t", n, s, neighbors, important, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	s := 0
	neighbors := [][]int{
		{1, 2},
		{0},
		{0},
	}
	important := []int{1, 2}
	runSample(t, n, s, neighbors, important, false)
}

func TestSample2(t *testing.T) {
	n := 3
	s := 0
	neighbors := [][]int{
		{1, 2},
		{0},
		{0},
	}
	important := []int{1}
	runSample(t, n, s, neighbors, important, true)
}

func TestSample3(t *testing.T) {
	n := 7
	s := 0
	neighbors := [][]int{
		{1, 2},
		{0, 3, 4},
		{0, 5, 6},
		{1},
		{1},
		{2},
		{2},
	}
	important := []int{3, 4, 6}
	runSample(t, n, s, neighbors, important, true)
}
