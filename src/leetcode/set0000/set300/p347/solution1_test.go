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
