package p2640

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, nums []int, expect []int64) {
	res := findPrefixScore(nums)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Samaple expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{2, 3, 7, 5, 10}
	expect := []int64{4, 10, 24, 36, 56}
	runSample(t, nums, expect)
}
