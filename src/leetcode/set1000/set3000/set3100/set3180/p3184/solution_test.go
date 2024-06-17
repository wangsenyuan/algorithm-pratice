package p3184

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, nums []int, queries [][]int, expect []int) {
	res := countOfPeaks(nums, queries)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{3, 1, 4, 2, 5}
	queries := [][]int{
		{2, 3, 4},
		{1, 0, 4},
	}
	expect := []int{0}
	runSample(t, nums, queries, expect)
}
func TestSample2(t *testing.T) {
	nums := []int{4, 1, 4, 2, 1, 5}
	queries := [][]int{
		{2, 2, 4},
		{1, 0, 2},
		{1, 0, 4},
	}
	expect := []int{0, 1}
	runSample(t, nums, queries, expect)
}

func TestSample3(t *testing.T) {
	nums := []int{3, 6, 9}
	queries := [][]int{
		{1, 1, 1},
		{1, 2, 2},
		{2, 2, 3},
	}
	expect := []int{0, 0}
	runSample(t, nums, queries, expect)
}
