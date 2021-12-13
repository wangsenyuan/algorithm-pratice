package p2089

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, nums []int, k int, expect []int) {
	res := getAverages(nums, k)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{7, 4, 3, 9, 1, 8, 5, 2, 6}
	k := 3
	expect := []int{-1, -1, -1, 5, 4, 4, -1, -1, -1}
	runSample(t, nums, k, expect)
}
