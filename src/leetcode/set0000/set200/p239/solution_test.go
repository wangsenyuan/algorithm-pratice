package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, nums []int, k int, expect []int) {
	res := maxSlidingWindow(nums, k)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{1, -1}
	k := 1
	expect := []int{1, -1}
	runSample(t, nums, k, expect)
}
