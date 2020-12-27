package p1706

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, nums []int, queries [][]int, expect []int) {
	res := maximizeXor(nums, queries)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{0, 1, 2, 3, 4}
	qs := [][]int{
		{3, 1}, {1, 3}, {5, 6},
	}
	expect := []int{3, 3, 7}
	runSample(t, nums, qs, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{5, 2, 4, 6, 6, 3}
	qs := [][]int{
		{12, 4}, {8, 1}, {6, 3},
	}
	expect := []int{15, -1, 5}
	runSample(t, nums, qs, expect)
}

func TestSample3(t *testing.T) {
	nums := []int{536870912, 0, 534710168, 330218644, 142254206}
	qs := [][]int{
		{558240772, 1000000000}, {307628050, 1000000000}, {3319300, 1000000000}, {2751604, 683297522}, {214004, 404207941},
	}
	expect := []int{1050219420, 844498962, 540190212, 539622516, 330170208}
	runSample(t, nums, qs, expect)
}
