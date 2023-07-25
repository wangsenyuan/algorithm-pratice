package main

import "testing"

func runSample(t *testing.T, n int, rocks [][]int, expect int) {
	res := solve(n, rocks)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	rocks := [][]int{
		{2, 3},
	}
	expect := 1
	runSample(t, n, rocks, expect)
}

func TestSample2(t *testing.T) {
	n := 3
	rocks := [][]int{
		{2, 1},
		{1, 2},
	}
	expect := 3
	runSample(t, n, rocks, expect)
}

func TestSample3(t *testing.T) {
	n := 5
	rocks := [][]int{
		{2, 3},
		{3, 1},
		{1, 2},
	}
	expect := 4
	runSample(t, n, rocks, expect)
}

func TestSample4(t *testing.T) {
	n := 5
	rocks := [][]int{
		{4, 5},
		{5, 1},
		{2, 2},
		{3, 3},
	}
	expect := 2
	runSample(t, n, rocks, expect)
}

func TestSample5(t *testing.T) {
	n := 2
	rocks := [][]int{
		{2, 2},
	}
	expect := 0
	runSample(t, n, rocks, expect)
}
