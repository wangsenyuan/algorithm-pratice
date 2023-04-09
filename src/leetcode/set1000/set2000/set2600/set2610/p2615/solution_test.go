package p2615

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, nums []int, expect []int64) {
	res := distance(nums)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{1, 3, 1, 1, 2}
	expect := []int64{5, 0, 3, 4, 0}
	runSample(t, nums, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{0, 5, 3}
	expect := []int64{0, 0, 0}
	runSample(t, nums, expect)
}
