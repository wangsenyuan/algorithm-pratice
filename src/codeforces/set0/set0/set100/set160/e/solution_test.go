package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, buses [][]int, people [][]int, expect []int) {
	res := solve(buses, people)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	buses := [][]int{
		{1, 10, 10},
		{5, 6, 2},
		{6, 7, 3},
		{5, 7, 4},
	}
	people := [][]int{
		{5, 7, 1},
		{1, 2, 1},
		{1, 10, 11},
	}
	expect := []int{4, 1, -1}
	runSample(t, buses, people, expect)
}

func TestSample2(t *testing.T) {
	buses := [][]int{
		{1, 1000000000, 1000000000},
	}
	people := [][]int{
		{1, 1000000000, 1000000000},
	}
	expect := []int{1}
	runSample(t, buses, people, expect)
}

func TestSample3(t *testing.T) {
	buses := [][]int{
		{2, 5, 986214363},
		{1, 2, 781201673},
		{1, 5, 644142615},
		{1, 3, 763531510},
		{3, 5, 820385107},
		{2, 3, 895170353},
		{2, 3, 359397940},
		{1, 3, 373910642},
		{2, 3, 493218144},
		{1, 2, 616153357},
	}
	people := [][]int{
		{1, 3, 3},
		{2, 5, 781201673},
		{3, 4, 1},
		{2, 3, 1},
		{3, 4, 4},
		{3, 5, 359397940},
		{1, 2, 5},
		{2, 5, 493218144},
		{2, 3, 493218144},
		{1, 3, 1},
	}
	expect := []int{8, 1, 3, 7, 3, 3, 8, 3, 9, 8}
	runSample(t, buses, people, expect)
}
