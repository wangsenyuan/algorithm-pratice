package p3275

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, queries [][]int, k int, expect []int) {
	res := resultsArray(queries, k)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}
func TestSample1(t *testing.T) {
	queries := [][]int{
		{1, 2},
		{3, 4},
		{2, 3},
		{-3, 0},
	}
	k := 2
	expect := []int{-1, 7, 5, 3}
	runSample(t, queries, k, expect)
}
