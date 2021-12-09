package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, k int, nums []int, expect []int) {
	res := solve(k, nums)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	k := 2
	nums := []int{1, 2, 4}
	expect := []int{0, 1, 166666669}
	runSample(t, k, nums, expect)
}
