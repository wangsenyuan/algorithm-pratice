package p2937

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, heights []int, queries [][]int, expect []int) {
	res := leftmostBuildingQueries(heights, queries)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	heights := []int{6, 4, 8, 5, 2, 7}
	queries := [][]int{{0, 1}, {0, 3}, {2, 4}, {3, 4}, {2, 2}}
	expect := []int{2, 5, -1, 5, 2}
	runSample(t, heights, queries, expect)
}

func TestSample2(t *testing.T) {
	heights := []int{5, 3, 8, 2, 6, 1, 4, 6}
	queries := [][]int{{0, 7}, {3, 5}, {5, 2}, {3, 0}, {1, 6}}
	expect := []int{7, 6, -1, 4, 6}
	runSample(t, heights, queries, expect)
}
