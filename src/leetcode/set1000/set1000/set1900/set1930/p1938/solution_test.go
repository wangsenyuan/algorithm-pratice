package p1938

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, parents []int, queries [][]int, expect []int) {
	res := maxGeneticDifference(parents, queries)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	parents := []int{-1, 0, 1, 1}
	queries := [][]int{
		{0, 2}, {3, 2}, {2, 5},
	}
	expect := []int{2, 3, 7}
	runSample(t, parents, queries, expect)
}

func TestSample2(t *testing.T) {
	parents := []int{3, 7, -1, 2, 0, 7, 0, 2}
	queries := [][]int{
		{4, 6}, {1, 15}, {0, 5},
	}
	expect := []int{6, 14, 7}
	runSample(t, parents, queries, expect)
}

func TestSample3(t *testing.T) {
	parents := []int{3, 5, 5, 7, 9, -1, 5, 5, 9, 2}
	queries := [][]int{
		{3, 34}, {2, 46}, {9, 36}, {7, 15}, {7, 26}, {3, 6}, {1, 42}, {6, 22}, {6, 26}, {7, 46},
	}
	expect := []int{39, 44, 45, 10, 31, 5, 47, 19, 31, 43}
	runSample(t, parents, queries, expect)
}
