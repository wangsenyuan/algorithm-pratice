package p3266

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, nums []int, k int, muliplier int, expect []int) {
	res := getFinalState(nums, k, muliplier)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{2, 1, 3, 5, 6}
	k := 5
	muliplier := 2
	expect := []int{8, 4, 6, 5, 6}
	runSample(t, nums, k, muliplier, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{100000, 2000}
	k := 2
	muliplier := 1000000
	expect := []int{999999307, 999999993}
	runSample(t, nums, k, muliplier, expect)
}

func TestSample3(t *testing.T) {
	nums := []int{2, 1, 3, 5, 6}
	k := 10000
	muliplier := 2
	expect := []int{996874050, 498437025, 247655534, 873046283, 247655534}
	runSample(t, nums, k, muliplier, expect)
}
