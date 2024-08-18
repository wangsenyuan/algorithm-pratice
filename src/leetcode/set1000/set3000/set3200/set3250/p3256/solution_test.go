package p3256

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, s string, k int, queries [][]int, expect []int64) {
	res := countKConstraintSubstrings(s, k, queries)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "0001111"
	k := 2
	queries := [][]int{
		{0, 6},
	}
	expect := []int64{26}
	runSample(t, s, k, queries, expect)
}

func TestSample2(t *testing.T) {
	s := "010101"
	k := 1
	queries := [][]int{
		{0, 5},
		{1, 4},
		{2, 3},
	}
	expect := []int64{15, 9, 3}
	runSample(t, s, k, queries, expect)
}
