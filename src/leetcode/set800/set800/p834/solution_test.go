package p834

import (
	"testing"
	"reflect"
)

func runSample(t *testing.T, N int, edges [][]int, expect []int) {
	res := sumOfDistancesInTree(N, edges)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("sample %d %v, expect %v, but got %v", N, edges, expect, res)
	}
}

func TestSample1(t *testing.T) {
	N := 6
	edges := [][]int{{0, 1}, {0, 2}, {2, 3}, {2, 4}, {2, 5}}
	expect := []int{8, 12, 6, 10, 10, 10}
	runSample(t, N, edges, expect)
}
