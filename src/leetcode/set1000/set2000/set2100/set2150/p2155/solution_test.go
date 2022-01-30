package p2155

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, nums []int, expect []int) {
	res := maxScoreIndices(nums)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{0, 0, 1, 0}
	expect := []int{2, 4}
	runSample(t, nums, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{0, 0, 0}
	expect := []int{3}
	runSample(t, nums, expect)
}

func TestSample3(t *testing.T) {
	nums := []int{1, 1, 1}
	expect := []int{0}
	runSample(t, nums, expect)
}
