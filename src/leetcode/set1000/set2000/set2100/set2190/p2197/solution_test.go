package p2197

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, nums []int, expect []int) {
	res := replaceNonCoprimes(nums)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{6, 4, 3, 2, 7, 6, 2}
	expect := []int{12, 7, 6}
	runSample(t, nums, expect)
}
