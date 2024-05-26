package p3159

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, limit int, queries [][]int, expect []int) {
	res := queryResults(limit, queries)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	limit := 4
	queries := [][]int{
		{1, 4}, {2, 5}, {1, 3}, {3, 4},
	}
	expect := []int{1, 2, 2, 3}
	runSample(t, limit, queries, expect)
}
