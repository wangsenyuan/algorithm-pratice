package p3425

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
		{0, 1, 2}, {1, 2, 3}, {1, 3, 5}, {1, 4, 4}, {2, 5, 6},
	}
	nums := []int{2, 1, 2, 1, 3, 1}
	expect := []int{6, 2}
	runSample(t, edges, nums, expect)
}

func TestSample2(t *testing.T) {
	edges := [][]int{
		{0, 1, 8},
	}
	nums := []int{2, 2}
	expect := []int{0, 1}
	runSample(t, edges, nums, expect)
}
