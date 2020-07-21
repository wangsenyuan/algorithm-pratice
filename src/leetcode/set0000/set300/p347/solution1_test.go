package p347

import (
	"reflect"
	"testing"
)

func TestTopKFrequent(t *testing.T) {
	nums := []int{3, 0, 1, 0}
	k := 1
	expect := []int{0}
	got := topKFrequent1(nums, k)
	if !reflect.DeepEqual(got, expect) {
		t.Errorf("expect %v; but get %v\n", expect, got)
	}
}

func TestTopKFrequent1(t *testing.T) {
	nums := []int{1, 1, 2, 2, 3, 3}
	k := 2
	expect := [][]int{
		[]int{1, 2},
		[]int{1, 3},
		[]int{2, 3},
	}
	got := topKFrequent1(nums, k)
	find := false
	for _, x := range expect {
		if reflect.DeepEqual(got, x) {
			find = true
			break
		}
	}

	if !find {
		t.Errorf("expect %v; but get %v\n", expect, got)
	}
}
