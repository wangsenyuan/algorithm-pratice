package main

import "testing"

func runSample(t *testing.T, x int, songs [][]int, expect int) {
	res := solve(x, songs)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	x := 3
	songs := [][]int{
		{1, 1},
		{1, 2},
		{1, 3},
	}
	expect := 6
	runSample(t, x, songs, expect)
}

func TestSample2(t *testing.T) {
	x := 10
	songs := [][]int{
		{5, 3},
		{2, 1},
		{3, 2},
		{5, 1},
	}
	expect := 10
	runSample(t, x, songs, expect)
}
