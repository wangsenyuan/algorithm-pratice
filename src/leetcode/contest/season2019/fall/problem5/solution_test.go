package problem5

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, leadership [][]int, operations [][]int, expect []int) {
	res := bonus(n, leadership, operations)
	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample %d %v %v, expect %v, but got %v", n, leadership, operations, expect, res)
	}
}

func TestSample1(t *testing.T) {
	N := 6
	leadership := [][]int{{1, 2}, {1, 6}, {2, 3}, {2, 5}, {1, 4}}
	operations := [][]int{{1, 1, 500}, {2, 2, 50}, {3, 1}, {2, 6, 15}, {3, 1}}
	expect := []int{650, 665}
	runSample(t, N, leadership, operations, expect)
}
