package p078

import (
	"reflect"
	"testing"
)

func TestSample(t *testing.T) {
	nums := []int{1, 2, 3}
	expect := [][]int{
		{3},
		{1},
		{2},
		{1, 2, 3},
		{1, 3},
		{2, 3},
		{1, 2},
		{},
	}
	res := subsets(nums)

	if len(res) != len(expect) {
		t.Errorf("sample expect %v, but got %v", expect, res)
	}

	for _, x := range expect {
		var found bool
		for _, y := range res {
			if reflect.DeepEqual(x, y) {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("expect %v, but not found", x)
		}
	}
}
