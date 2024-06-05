package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, a, b int, clips [][]int, operations []string, expect []int) {
	res := solve(a, b, clips, operations)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a, b := 4, 4
	clips := [][]int{
		{4, 1},
		{3, 3},
		{2, 4},
	}
	ops := []string{
		"D 2",
		"R 1",
	}
	expect := []int{2, 1}
	runSample(t, a, b, clips, ops, expect)
}

func TestSample2(t *testing.T) {
	a, b := 4, 4
	clips := [][]int{
		{4, 1},
		{3, 2},
		{2, 3},
	}
	ops := []string{
		"D 1",
		"L 1",
		"U 2",
	}
	expect := []int{2, 0}
	runSample(t, a, b, clips, ops, expect)
}

func TestSample3(t *testing.T) {
	a, b := 3, 5
	clips := [][]int{
		{1, 3},
		{2, 2},
		{3, 3},
	}
	ops := []string{
		"R 2",
		"R 2",
	}
	expect := []int{0, 3}
	runSample(t, a, b, clips, ops, expect)
}
