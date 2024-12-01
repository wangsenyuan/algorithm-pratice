package p3370

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, edges1 [][]int, edges2 [][]int, expect []int) {
	res := maxTargetNodes(edges1, edges2)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	edges1 := [][]int{
		{0, 1}, {0, 2}, {2, 3}, {2, 4},
	}
	edges2 := [][]int{
		{0, 1}, {0, 2}, {0, 3}, {2, 7}, {1, 4}, {4, 5}, {4, 6},
	}
	expect := []int{8, 7, 7, 8, 8}
	runSample(t, edges1, edges2, expect)
}

func TestSample2(t *testing.T) {
	edges1 := [][]int{
		{0, 1}, {0, 2}, {0, 3}, {0, 4},
	}
	edges2 := [][]int{
		{0, 1}, {1, 2}, {2, 3},
	}
	expect := []int{3, 6, 6, 6, 6}
	runSample(t, edges1, edges2, expect)
}
