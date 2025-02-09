package p3447

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, groups []int, elements []int, expect []int) {
	res := assignElements(groups, elements)
	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample %v %v, expect %v, but got %v", groups, elements, expect, res)
	}
}

func TestSample1(t *testing.T) {
	groups := []int{8, 4, 3, 2, 4}
	elements := []int{4, 2}
	expect := []int{0, 0, -1, 1, 0}
	runSample(t, groups, elements, expect)
}

func TestSample2(t *testing.T) {
	groups := []int{2, 3, 5, 7}
	elements := []int{5, 3, 3}
	expect := []int{-1, 1, 0, -1}
	runSample(t, groups, elements, expect)
}
