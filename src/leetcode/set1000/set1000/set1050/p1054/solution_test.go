package p1054

import (
	"reflect"
	"sort"
	"testing"
)

func runSample(t *testing.T, nums []int) {
	res := rearrangeBarcodes(nums)

	for i := 1; i < len(nums); i++ {
		if res[i] == res[i-1] {
			t.Errorf("Sample %v, result %v is wrong", nums, res)
		}
	}

	sort.Ints(nums)
	sort.Ints(res)

	if !reflect.DeepEqual(res, nums) {
		t.Errorf("Sample %v, result %v is wrong", nums, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{1, 1, 1, 2, 2, 2}
	runSample(t, nums)
}

func TestSample2(t *testing.T) {
	nums := []int{1, 1, 1, 1, 2, 2, 3, 3}
	runSample(t, nums)
}

func TestSample3(t *testing.T) {
	nums := []int{7, 7, 7, 8, 5, 7, 5, 5, 5, 8}
	runSample(t, nums)
}
