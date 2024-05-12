package p3142

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, nums []int, expect []int) {
	res := findPermutation(nums)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v but got %v", expect, res)
	}
}

func calc(nums []int, perm []int) int {
	var res int
	n := len(nums)

	for i := 0; i+1 < n; i++ {
		res += abs(perm[i] - nums[perm[i+1]])
	}

	res += abs(perm[n-1] - nums[perm[0]])

	return res
}

func TestSample1(t *testing.T) {
	nums := []int{1, 0, 2}
	expect := []int{0, 1, 2}
	runSample(t, nums, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{0, 2, 1}
	expect := []int{0, 2, 1}
	runSample(t, nums, expect)
}

func TestSample3(t *testing.T) {
	nums := []int{1, 0, 2, 4, 5, 3}
	expect := []int{0, 1, 2, 5, 4, 3}
	runSample(t, nums, expect)
}

func TestSample4(t *testing.T) {
	nums := []int{0, 3, 2, 1}
	expect := []int{0, 2, 3, 1}
	runSample(t, nums, expect)
}
