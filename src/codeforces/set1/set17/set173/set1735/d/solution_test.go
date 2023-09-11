package main

import "testing"

func runSample(t *testing.T, cards [][]int, expect int) {
	res := int(solve(cards))

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	cards := [][]int{
		{0, 0},
		{0, 1},
		{0, 2},
		{1, 0},
		{1, 1},
		{1, 2},
		{2, 0},
		{2, 1},
		{2, 2},
	}
	expect := 54
	runSample(t, cards, expect)
}
