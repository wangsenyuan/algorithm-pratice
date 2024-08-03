package main

import "testing"

func runSample(t *testing.T, items [][]int, expect int) {
	res := solve(items)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	items := [][]int{
		{2, 10},
		{0, 20},
		{1, 5},
		{1, 3},
	}
	expect := 8
	runSample(t, items, expect)
}

func TestSample2(t *testing.T) {
	items := [][]int{
		{0, 1},
		{0, 10},
		{0, 100},
	}
	expect := 111
	runSample(t, items, expect)
}

func TestSample3(t *testing.T) {
	items := [][]int{
		{2, 87623264},
		{0, 864627704},
	}
	expect := 87623264
	runSample(t, items, expect)
}
