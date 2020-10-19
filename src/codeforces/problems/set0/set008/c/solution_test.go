package main

import "testing"

func runSample(t *testing.T, home []int, items [][]int, expect int) {
	res, path := solve(home, items)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}

	var tot int

	for i := 1; i < len(path); i++ {
		if path[i] != 0 {
			if path[i-1] == 0 {
				tot += distance(home, items[path[i]-1])
			} else {
				tot += distance(items[path[i-1]-1], items[path[i]-1])
			}
		} else {
			tot += distance(items[path[i-1]-1], home)
		}
	}
	if tot != expect {
		t.Fatalf("Sample result %v, distance %d not expected %d", res, tot, expect)
	}
}

func TestSample1(t *testing.T) {
	home := []int{0, 0}
	items := [][]int{
		{1, 1},
		{-1, 1},
	}
	expect := 8
	runSample(t, home, items, expect)
}

func TestSample2(t *testing.T) {
	home := []int{1, 1}
	items := [][]int{
		{4, 3},
		{3, 4},
		{0, 0},
	}
	expect := 32
	runSample(t, home, items, expect)
}

func TestSample3(t *testing.T) {
	home := []int{-44, 47}
	items := [][]int{
		{96, -18},
		{-50, 86},
		{84, 68},
		{-25, 80},
		{-11, -15},
		{-62, 0},
		{-42, 50},
		{-57, 11},
		{-5, 27},
		{-44, 67},
		{-77, -3},
		{-27, -46},
		{32, 63},
		{86, 13},
		{-21, -51},
		{-25, -62},
		{-14, -2},
		{-21, 86},
		{-92, -94},
		{-44, -34},
		{-74, 55},
		{91, -35},
		{-10, 55},
		{-34, 16},
	}
	expect := 191534

	runSample(t, home, items, expect)
}
