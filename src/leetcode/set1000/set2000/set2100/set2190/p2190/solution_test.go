package p2190

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, mapping []int, nums []int, expect []int) {
	res := sortJumbled(mapping, nums)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	mapping := []int{8, 9, 4, 0, 2, 1, 3, 5, 7, 6}
	nums := []int{991, 338, 38}
	expect := []int{338, 38, 991}
	runSample(t, mapping, nums, expect)
}

func TestSample2(t *testing.T) {
	mapping := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	nums := []int{789, 456, 123}
	expect := []int{123, 456, 789}
	runSample(t, mapping, nums, expect)
}

func TestSample3(t *testing.T) {
	mapping := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	nums := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	expect := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	runSample(t, mapping, nums, expect)
}
