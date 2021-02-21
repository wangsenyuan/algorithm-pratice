package p1765

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, nums []int, edges [][]int, expect []int) {
	res := getCoprimes(nums, edges)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{2, 3, 3, 2}
	edges := [][]int{
		{0, 1}, {1, 2}, {1, 3},
	}
	expect := []int{-1, 0, 0, 1}
	runSample(t, nums, edges, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{5, 6, 10, 2, 3, 6, 15}
	edges := [][]int{
		{0, 1}, {0, 2}, {1, 3}, {1, 4}, {2, 5}, {2, 6},
	}
	expect := []int{-1, 0, -1, 0, 0, 0, -1}
	runSample(t, nums, edges, expect)
}
