package c

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, increase [][]int, requirements [][]int, expect []int) {
	res := getTriggerTime(increase, requirements)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	incr := [][]int{{0, 4, 5}, {4, 8, 8}, {8, 6, 1}, {10, 10, 0}}
	req := [][]int{{12, 11, 16}, {20, 2, 6}, {9, 2, 6}, {10, 18, 3}, {8, 14, 9}}
	expect := []int{-1, 4, 3, 3, 3}
	runSample(t, incr, req, expect)
}

func TestSample2(t *testing.T) {
	incr := [][]int{{2, 8, 4}, {2, 5, 0}, {10, 9, 8}}
	req := [][]int{{2, 11, 3}, {15, 10, 7}, {9, 17, 12}, {8, 1, 14}}
	expect := []int{2, -1, 3, -1}
	runSample(t, incr, req, expect)
}
