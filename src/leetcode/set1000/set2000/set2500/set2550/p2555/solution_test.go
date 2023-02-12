package p2555

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, s string, queries [][]int, expect [][]int) {
	res := substringXorQueries(s, queries)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "101101"
	queries := [][]int{
		{0, 5},
		{1, 2},
	}
	expect := [][]int{
		{0, 2},
		{2, 3},
	}
	runSample(t, s, queries, expect)
}

func TestSample2(t *testing.T) {
	s := "0101"
	queries := [][]int{
		{12, 8},
	}
	expect := [][]int{
		{-1, -1},
	}
	runSample(t, s, queries, expect)
}
