package p2859

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, edges [][]int, expect []int) {
	res := minEdgeReversals(n, edges)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 4
	edges := [][]int{
		{2, 0}, {2, 1}, {1, 3},
	}
	expect := []int{1, 1, 0, 2}
	runSample(t, n, edges, expect)
}

func TestSample2(t *testing.T) {
	n := 3
	edges := [][]int{
		{1, 2}, {2, 0},
	}
	expect := []int{2, 0, 1}
	runSample(t, n, edges, expect)
}
