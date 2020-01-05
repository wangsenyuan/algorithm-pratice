package p1310

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, arr []int, queries [][]int, expect []int) {
	res := xorQueries(arr, queries)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample %v %v, expect %v, but got %v ", arr, queries, expect, res)
	}
}

func TestSample1(t *testing.T) {
	arr := []int{1, 3, 4, 8}
	queries := [][]int{{0, 1}, {1, 2}, {0, 3}, {3, 3}}
	expect := []int{2, 7, 14, 8}
	runSample(t, arr, queries, expect)
}
