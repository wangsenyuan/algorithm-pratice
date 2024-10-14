package main

import "testing"

func runSample(t *testing.T, k int, castles [][]int, portals [][]int, expect int) {
	res := solve(k, castles, portals)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	k := 7
	castles := [][]int{
		{7, 4, 17},
		{3, 0, 8},
		{11, 2, 0},
		{13, 3, 5},
	}
	portals := [][]int{
		{3, 1},
		{2, 1},
		{4, 3},
	}
	expect := 5
	runSample(t, k, castles, portals, expect)
}

func TestSample2(t *testing.T) {
	k := 7
	castles := [][]int{
		{7, 4, 17},
		{3, 0, 8},
		{11, 2, 0},
		{13, 3, 5},
	}
	portals := [][]int{
		{3, 1},
		{2, 1},
		{4, 1},
	}
	expect := 22
	runSample(t, k, castles, portals, expect)
}

func TestSample3(t *testing.T) {
	k := 7
	castles := [][]int{
		{7, 4, 17},
		{3, 0, 8},
		{11, 2, 0},
		{14, 3, 5},
	}
	portals := [][]int{
		{3, 1},
		{2, 1},
		{4, 3},
	}
	expect := -1
	runSample(t, k, castles, portals, expect)
}

func TestSample4(t *testing.T) {
	k := 2
	castles := [][]int{
		{2, 1, 1615},
		{2, 1, 1326},
		{0, 0, 3880},
		{2, 0, 1596},
		{1, 2, 4118},
		{0, 0, 4795},
		{0, 1, 814},
		{0, 0, 4555},
		{0, 1, 3284},
		{2, 2, 1910},
	}
	portals := [][]int{
		{10, 6},
		{2, 1},
		{8, 3},
		{8, 7},
		{6, 3},
		{4, 1},
		{7, 4},
		{6, 4},
		{7, 2},
		{10, 9},
	}
	expect := 27079
	runSample(t, k, castles, portals, expect)
}
