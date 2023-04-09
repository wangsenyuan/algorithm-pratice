package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, E [][]int, weight int, sum int, expect []int) {
	w, s, res := solve(n, E)

	if w != weight || s != sum || !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %d %d %v, but got %d %d %v", weight, sum, expect, w, s, res)
	}
}

func TestSample1(t *testing.T) {
	n := 4
	E := [][]int{
		{1, 2},
		{2, 3},
		{2, 4},
	}
	weight := 3
	sum := 4
	expect := []int{1, 1, 1, 1}
	runSample(t, n, E, weight, sum, expect)
}

func TestSample2(t *testing.T) {
	n := 9
	E := [][]int{
		{3, 4},
		{7, 6},
		{2, 1},
		{8, 3},
		{5, 6},
		{1, 8},
		{8, 6},
		{9, 6},
	}
	weight := 6
	sum := 11
	expect := []int{1, 1, 1, 1, 1, 1, 1, 3, 1}
	runSample(t, n, E, weight, sum, expect)
}

func TestSample3(t *testing.T) {
	n := 4
	E := [][]int{
		{1, 2},
		{2, 3},
		{3, 4},
	}
	weight := 2
	sum := 4
	expect := []int{1, 1, 1, 1}
	runSample(t, n, E, weight, sum, expect)
}
