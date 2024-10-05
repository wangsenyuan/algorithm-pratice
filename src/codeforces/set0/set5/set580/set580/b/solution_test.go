package main

import "testing"

func runSample(t *testing.T, friends [][]int, d int, expect int) {
	res := solve(friends, d)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}
func TestSample1(t *testing.T) {
	d := 5
	friends := [][]int{
		{75, 5},
		{0, 100},
		{150, 20},
		{75, 1},
	}
	expect := 100
	runSample(t, friends, d, expect)
}

func TestSample2(t *testing.T) {
	d := 2
	friends := [][]int{
		{10909234, 9},
		{10909236, 8},
		{10909237, 10},
		{10909235, 98},
	}
	expect := 107
	runSample(t, friends, d, expect)
}
