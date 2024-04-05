package main

import "testing"

func runSample(t *testing.T, people [][]int, x int, y int, z int, expect int64) {
	res := solve(people, x, y, z)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	x, y, z := 1, 2, 1
	people := [][]int{
		{2, 4, 4},
		{3, 2, 1},
		{7, 6, 7},
		{5, 2, 3},
	}
	var expect int64 = 18
	runSample(t, people, x, y, z, expect)
}

func TestSample2(t *testing.T) {
	x, y, z := 3, 3, 2
	people := [][]int{
		{16, 17, 1},
		{2, 7, 5},
		{2, 16, 12},
		{17, 7, 7},
		{13, 2, 10},
		{12, 18, 3},
		{16, 15, 19},
		{5, 6, 2},
	}
	var expect int64 = 110
	runSample(t, people, x, y, z, expect)
}

func TestSample3(t *testing.T) {
	x, y, z := 6, 2, 4
	people := [][]int{
		{33189, 87907, 277349742},
		{71616, 46764, 575306520},
		{8801, 53151, 327161251},
		{58589, 4337, 796697686},
		{66854, 17565, 289910583},
		{50598, 35195, 478112689},
		{13919, 88414, 103962455},
		{7953, 69657, 699253752},
		{44255, 98144, 468443709},
		{2332, 42580, 752437097},
		{39752, 19060, 845062869},
		{60126, 74101, 382963164},
	}
	var expect int64 = 3093929975
	runSample(t, people, x, y, z, expect)
}
