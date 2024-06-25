package main

import "testing"

func runSample(t *testing.T, disks [][]int, expect bool) {
	res := solve(disks)

	if res != expect {
		t.Fatalf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	disks := [][]int{
		{0, 2, 1},
		{0, 0, 1},
		{4, -3, 4},
		{11, 0, 3},
		{11, 5, 2},
	}
	expect := true
	runSample(t, disks, expect)
}

func TestSample2(t *testing.T) {
	disks := [][]int{
		{2, 2, 2},
		{7, 2, 3},
		{7, 7, 2},
		{2, 7, 3},
	}
	expect := false
	runSample(t, disks, expect)
}
