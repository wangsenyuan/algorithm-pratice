package p1926

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, nums []int, expect []int) {
	res := findMaximums(nums)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{0, 1, 2, 4}
	expect := []int{4, 2, 1, 0}
	runSample(t, nums, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{10, 20, 50, 10}
	expect := []int{50, 20, 10, 10}
	runSample(t, nums, expect)
}
