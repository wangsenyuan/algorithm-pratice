package p310

import (
	"reflect"
	"sort"
	"testing"
)

func runSample(t *testing.T, n int, edges [][]int, expect []int) {
	res := findMinHeightTrees(n, edges)

	sort.Ints(expect)
	sort.Ints(res)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 4
	edges := [][]int{
		{1, 0}, {1, 2}, {1, 3},
	}
	expect := []int{1}
	runSample(t, n, edges, expect)
}

func TestSample2(t *testing.T) {
	n := 6
	edges := [][]int{
		{3, 0}, {3, 1}, {3, 2}, {3, 4}, {5, 4},
	}
	expect := []int{3, 4}
	runSample(t, n, edges, expect)
}

func TestSample3(t *testing.T) {
	n := 10
	edges := [][]int{
		{0, 3}, {1, 3}, {2, 3}, {4, 3}, {5, 3}, {4, 6}, {4, 7}, {5, 8}, {5, 9},
	}
	expect := []int{3}
	runSample(t, n, edges, expect)
}
