package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, nums []int, expect []int64) {
	res := solve(nums)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{0, 1, 3}
	expect := []int64{1, 1, 0, -1}
	runSample(t, nums, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{0, 1, 2, 3, 4, 3, 2}
	expect := []int64{1, 1, 2, 2, 1, 0, 2, 6}
	runSample(t, nums, expect)
}

func TestSample3(t *testing.T) {
	nums := []int{3, 0, 0, 0}
	expect := []int64{3, 0, 1, 4, 3}
	runSample(t, nums, expect)
}

func TestSample4(t *testing.T) {
	nums := []int{4, 6, 2, 3, 5, 0, 5}
	expect := []int64{1, 0, -1, -1, -1, -1, -1, -1}
	runSample(t, nums, expect)
}

func TestSample5(t *testing.T) {
	nums := []int{4, 0, 1, 0, 4}
	expect := []int64{2, 1, 0, 2, -1, -1}
	runSample(t, nums, expect)
}
