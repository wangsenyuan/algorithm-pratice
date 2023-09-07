package main

import (
	"testing"
)

func runSample(t *testing.T, boxes [][]int, expect bool) {
	res := solve(boxes)

	if res != expect {
		t.Fatalf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	boxes := [][]int{
		{19, 8, 22},
		{10, 24, 12},
		{15, 25, 11},
	}
	expect := true
	runSample(t, boxes, expect)
}

func TestSample2(t *testing.T) {
	boxes := [][]int{
		{19, 8, 22},
		{10, 25, 12},
		{15, 24, 11},
	}
	expect := false
	runSample(t, boxes, expect)
}

func TestSample3(t *testing.T) {
	boxes := [][]int{
		{1, 1, 2},
		{1, 2, 2},
	}
	expect := false
	runSample(t, boxes, expect)
}
