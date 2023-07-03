package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, events [][]int, expect []Ans) {
	res := solve(n, events)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 10
	events := [][]int{
		{3, 5},
		{1, 3, 8, 3, 1},
		{3, 5},
		{1, 5, 10, -8, 2},
		{3, 5},
		{3, 10},
		{2, 1, 10, 0, 1},
		{3, 6},
		{2, 5, 7, 2, 1},
		{3, 6},
	}
	expect := []Ans{
		{false, -1},
		{true, 7},
		{true, 7},
		{true, -38},
		{true, 11},
		{true, 14},
	}
	runSample(t, n, events, expect)
}
