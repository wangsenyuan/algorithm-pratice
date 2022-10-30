package p2453

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, nums []int, expect []int) {
	res := secondGreaterElement(nums)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{2, 4, 0, 9, 6}
	expect := []int{9, 6, 6, -1, -1}
	runSample(t, nums, expect)
}
