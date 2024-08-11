package main

import "testing"

func runSample(t *testing.T, n int, chords [][]int, expect int) {
	res := solve(n, chords)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 4
	chords := [][]int{
		{8, 2},
		{1, 5},
	}
	expect := 4
	runSample(t, n, chords, expect)
}

func TestSample2(t *testing.T) {
	n := 10
	chords := [][]int{
		{14, 6},
		{2, 20},
		{9, 10},
		{13, 18},
		{15, 12},
		{11, 7},
	}
	expect := 14
	runSample(t, n, chords, expect)
}
