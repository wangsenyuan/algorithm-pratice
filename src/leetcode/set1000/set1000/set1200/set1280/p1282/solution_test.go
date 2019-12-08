package p1282

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, groupSizes []int, expect [][]int) {
	res := groupThePeople(groupSizes)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample %v, expect %v, but got %v", groupSizes, expect, res)
	}

}

func TestSample1(t *testing.T) {
	groupSizes := []int{3, 3, 3, 3, 3, 1, 3}
	expect := [][]int{{5}, {0, 1, 2}, {3, 4, 6}}
	runSample(t, groupSizes, expect)
}
