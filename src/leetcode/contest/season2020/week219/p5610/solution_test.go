package p5610

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, nums []int, expect []int) {
	res := getSumAbsoluteDifferences(nums)

	if !reflect.DeepEqual(expect, res) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{2, 3, 5}
	expect := []int{4, 3, 5}
	runSample(t, nums, expect)
}
