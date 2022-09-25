package p2420

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, nums []int, k int, expect []int) {
	res := goodIndices(nums, k)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{2, 1, 1, 1, 3, 4, 1}
	k := 2
	expect := []int{2, 3}
	runSample(t, nums, k, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{2, 1, 1, 2}
	k := 2
	runSample(t, nums, k, nil)
}

func TestSample3(t *testing.T) {
	nums := []int{243655, 573257, 92104, 932396, 605975, 645879, 104190, 32569}
	k := 3
	runSample(t, nums, k, nil)
}

func TestSample4(t *testing.T) {
	nums := []int{388589, 17165, 726687, 401298, 600033, 537254, 301052, 151069, 399955}
	k := 4
	runSample(t, nums, k, nil)
}
