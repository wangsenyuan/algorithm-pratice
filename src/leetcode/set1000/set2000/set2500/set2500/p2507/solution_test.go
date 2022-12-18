package p2507

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, queries [][]int, expect []int) {
	res := cycleLengthQueries(n, queries)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	queries := [][]int{
		{5, 3}, {4, 7}, {2, 3},
	}
	expect := []int{4, 5, 3}
	runSample(t, n, queries, expect)
}

func TestSample2(t *testing.T) {
	n := 2
	queries := [][]int{
		{1, 2},
	}
	expect := []int{2}
	runSample(t, n, queries, expect)
}
