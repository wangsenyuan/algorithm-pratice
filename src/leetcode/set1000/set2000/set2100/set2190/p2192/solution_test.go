package p2192

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, edges [][]int, expect [][]int) {
	res := getAncestors(n, edges)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 8
	edges := [][]int{
		{0, 3}, {0, 4}, {1, 3}, {2, 4}, {2, 7}, {3, 5}, {3, 6}, {3, 7}, {4, 6},
	}
	expect := [][]int{
		{}, {}, {}, {0, 1}, {0, 2}, {0, 1, 3}, {0, 1, 2, 3, 4}, {0, 1, 2, 3},
	}
	runSample(t, n, edges, expect)
}
