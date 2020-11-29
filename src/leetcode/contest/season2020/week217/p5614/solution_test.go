package p5614

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, nums []int, k int, expect []int) {
	res := mostCompetitive(nums, k)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{3, 5, 2, 6}
	k := 2
	expect := []int{2, 6}
	runSample(t, nums, k, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{2, 4, 3, 3, 5, 4, 9, 6}
	k := 4
	expect := []int{2, 3, 3, 4}
	runSample(t, nums, k, expect)
}

func TestSample3(t *testing.T) {
	nums := []int{371, 18, 52, 29, 55, 73, 24, 42, 66, 8, 80, 2}
	k := 3
	expect := []int{8, 80, 2}
	runSample(t, nums, k, expect)
}
