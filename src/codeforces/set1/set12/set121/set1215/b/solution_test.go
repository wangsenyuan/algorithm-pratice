package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, nums []int, expect []int64) {
	res := solve(nums)

	if !reflect.DeepEqual(expect, res) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{5, -3, 3, -1, 1}
	expect := []int64{8, 7}
	runSample(t, nums, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{4, 2, -4, 3, 1, 2, -4, 3, 2, 3}
	expect := []int64{28, 27}
	runSample(t, nums, expect)
}

func TestSample3(t *testing.T) {
	nums := []int{-1, -2, -3, -4, -5}
	expect := []int64{9, 6}
	runSample(t, nums, expect)
}
