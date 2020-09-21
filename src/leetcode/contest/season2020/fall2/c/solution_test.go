package c

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, nums []int, expect []int) {
	res := numsGame(nums)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample %v, expect %v, but got %v", nums, expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{3, 4, 5, 1, 6, 7}
	res := []int{0, 0, 0, 5, 6, 7}
	runSample(t, nums, res)
}

func TestSample2(t *testing.T) {
	nums := []int{1, 1, 1, 2, 3, 4}
	res := []int{0, 1, 2, 3, 3, 3}
	runSample(t, nums, res)
}
