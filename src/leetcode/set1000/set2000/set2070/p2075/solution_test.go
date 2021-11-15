package p2075

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, items [][]int, queries []int, expect []int) {
	res := maximumBeauty(items, queries)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	items := [][]int{
		{1, 2}, {3, 2}, {2, 4}, {5, 6}, {3, 5},
	}
	queries := []int{1, 2, 3, 4, 5, 6}
	expect := []int{2, 4, 5, 5, 6, 6}
	runSample(t, items, queries, expect)
}
