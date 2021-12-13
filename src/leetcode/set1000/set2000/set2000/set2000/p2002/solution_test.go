package p2002

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, parents []int, nums []int, expect []int) {
	res := smallestMissingValueSubtree(parents, nums)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	parents := []int{-1, 0, 0, 2}
	nums := []int{1, 2, 3, 4}
	expect := []int{5, 1, 1, 1}
	runSample(t, parents, nums, expect)
}

func TestSample2(t *testing.T) {
	parents := []int{-1, 0, 1, 0, 3, 3}
	nums := []int{5, 4, 6, 2, 1, 3}
	expect := []int{7, 1, 1, 4, 2, 1}
	runSample(t, parents, nums, expect)
}

func TestSample3(t *testing.T) {
	parents := []int{-1, 2, 3, 0, 2, 4, 1}
	nums := []int{2, 3, 4, 5, 6, 7, 8}
	expect := []int{1, 1, 1, 1, 1, 1, 1}
	runSample(t, parents, nums, expect)
}
