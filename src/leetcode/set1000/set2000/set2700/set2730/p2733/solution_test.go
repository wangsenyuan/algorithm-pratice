package p2733

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, nums1 []int, nums2 []int, queries [][]int, expect []int) {
	res := maximumSumQueries(nums1, nums2, queries)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums1 := []int{4, 3, 1, 2}
	nums2 := []int{2, 4, 9, 5}
	queries := [][]int{
		{4, 1}, {1, 3}, {2, 5},
	}
	expect := []int{6, 10, 7}
	runSample(t, nums1, nums2, queries, expect)
}

func TestSample2(t *testing.T) {
	nums1 := []int{3, 2, 5}
	nums2 := []int{2, 3, 4}
	queries := [][]int{
		{4, 4}, {3, 2}, {1, 1},
	}
	expect := []int{9, 9, 9}
	runSample(t, nums1, nums2, queries, expect)
}

func TestSample3(t *testing.T) {
	nums1 := []int{2, 1}
	nums2 := []int{2, 3}
	queries := [][]int{
		{3, 3},
	}
	expect := []int{-1}
	runSample(t, nums1, nums2, queries, expect)
}

func TestSample4(t *testing.T) {
	nums1 := []int{31}
	nums2 := []int{17}
	queries := [][]int{
		{1, 79},
	}
	expect := []int{-1}
	runSample(t, nums1, nums2, queries, expect)
}
