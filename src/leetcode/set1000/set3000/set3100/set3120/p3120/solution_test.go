package p3120

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, edges [][]int, expect []bool) {
	res := findAnswer(n, edges)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 6
	edges := [][]int{
		{0, 1, 4}, {0, 2, 1}, {1, 3, 2}, {1, 4, 3}, {1, 5, 1}, {2, 3, 1}, {3, 5, 3}, {4, 5, 2},
	}
	expect := []bool{true, true, true, false, true, true, true, false}
	runSample(t, n, edges, expect)
}
