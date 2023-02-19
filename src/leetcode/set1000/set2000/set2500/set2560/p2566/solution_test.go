package p2566

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, num1, num2 []int, qs [][]int, expect []int64) {
	res := handleQuery(num1, num2, qs)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums1 := []int{1, 0, 1}
	nums2 := []int{0, 0, 0}
	queries := [][]int{
		{1, 1, 1},
		{2, 1, 0},
		{3, 0, 0},
	}
	expect := []int64{3}
	runSample(t, nums1, nums2, queries, expect)
}

func TestSample2(t *testing.T) {
	nums1 := []int{1}
	nums2 := []int{5}
	queries := [][]int{
		{2, 0, 0},
		{3, 0, 0},
	}
	expect := []int64{5}
	runSample(t, nums1, nums2, queries, expect)
}

func TestSample3(t *testing.T) {
	nums1 := []int{0, 0, 0, 0, 1, 0, 1, 1, 1}
	nums2 := []int{35, 29, 21, 34, 8, 48, 22, 43, 37}
	queries := [][]int{
		{1, 4, 7},
		{3, 0, 0},
		{2, 27, 0},
		{3, 0, 0},
		{1, 0, 3},
		{3, 0, 0},
		{2, 6, 0},
		{1, 3, 8},
		{2, 13, 0},
		{3, 0, 0},
		{3, 0, 0},
		{3, 0, 0},
		{2, 2, 0},
		{2, 28, 0},
		{3, 0, 0},
		{3, 0, 0},
		{2, 25, 0},
		{3, 0, 0},
		{3, 0, 0},
		{1, 2, 5},
	}
	expect := []int64{277, 331, 331, 445, 445, 445, 625, 625, 775, 775}
	runSample(t, nums1, nums2, queries, expect)
}

func TestSample4(t *testing.T) {
	nums1 := []int{1, 0, 1}
	nums2 := []int{44, 28, 35}
	queries := [][]int{
		{1, 0, 1},
		{2, 10, 0},
		{2, 2, 0},
		{2, 7, 0},
		{3, 0, 0},
		{3, 0, 0},
		{1, 2, 2},
		{1, 1, 2},
		{2, 1, 0},
		{1, 0, 2},
		{1, 2, 2},
		{1, 0, 2},
		{3, 0, 0},
		{1, 1, 2},
		{3, 0, 0},
		{1, 0, 1},
		{2, 21, 0},
		{1, 0, 1},
		{2, 26, 0},
		{1, 1, 1},
	}
	expect := []int64{145, 145, 146, 146}
	runSample(t, nums1, nums2, queries, expect)
}
