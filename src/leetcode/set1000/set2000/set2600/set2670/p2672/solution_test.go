package p2672

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, queries [][]int, expect []int) {
	res := colorTheArray(n, queries)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 4
	queries := [][]int{
		{0, 2}, {1, 2}, {3, 1}, {1, 1}, {2, 1},
	}
	expect := []int{0, 1, 1, 0, 2}
	runSample(t, n, queries, expect)
}
