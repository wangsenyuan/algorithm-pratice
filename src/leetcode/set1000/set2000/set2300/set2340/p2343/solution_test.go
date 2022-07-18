package p2343

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, nums []string, queries [][]int, expect []int) {
	res := smallestTrimmedNumbers(nums, queries)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []string{"102", "473", "251", "814"}
	queries := [][]int{
		{1, 1}, {2, 3}, {4, 2}, {1, 2},
	}
	expect := []int{2, 2, 1, 0}
	runSample(t, nums, queries, expect)
}

func TestSample2(t *testing.T) {
	nums := []string{"24", "37", "96", "04"}
	queries := [][]int{
		{2, 1}, {2, 2},
	}
	expect := []int{3, 0}
	runSample(t, nums, queries, expect)
}
