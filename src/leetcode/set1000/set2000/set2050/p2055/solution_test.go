package p2055

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, s string, queries [][]int, expect []int) {
	res := platesBetweenCandles(s, queries)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "**|**|***|"
	queries := [][]int{
		{2, 5}, {5, 9},
	}
	expect := []int{2, 3}
	runSample(t, s, queries, expect)
}

func TestSample2(t *testing.T) {
	s := "***|**|*****|**||**|*"
	queries := [][]int{
		{1, 17}, {4, 5}, {14, 17}, {5, 11}, {15, 16},
	}
	expect := []int{9, 0, 0, 0, 0}
	runSample(t, s, queries, expect)
}
