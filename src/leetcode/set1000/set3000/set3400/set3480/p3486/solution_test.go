package p3486

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, edges [][]int, nums []int, expect []int) {
	res := longestSpecialPath(edges, nums)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	edges := [][]int{
		{0, 1, 1}, {1, 2, 3}, {1, 3, 1}, {2, 4, 6}, {4, 7, 2}, {3, 5, 2}, {3, 6, 5}, {6, 8, 3},
	}
	nums := []int{1, 1, 0, 3, 1, 2, 1, 1, 0}
	expect := []int{9, 3}
	runSample(t, edges, nums, expect)
}

func TestSample2(t *testing.T) {
	edges := [][]int{
		{1, 0, 3}, {0, 2, 4}, {0, 3, 5},
	}
	nums := []int{1, 1, 0, 2}
	expect := []int{5, 2}
	runSample(t, edges, nums, expect)
}

func TestSample3(t *testing.T) {
	edges := [][]int{
		{1, 0, 5}, {2, 1, 3},
	}
	nums := []int{3, 1, 1}
	expect := []int{8, 3}
	runSample(t, edges, nums, expect)
}
