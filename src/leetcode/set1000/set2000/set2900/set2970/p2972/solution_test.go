package p2972

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, edges [][]int, cost []int, expect []int64) {
	res := placedCoins(edges, cost)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	edges := [][]int{{0, 1}, {0, 2}, {0, 3}, {0, 4}, {0, 5}}
	cost := []int{1, 2, 3, 4, 5, 6}
	expect := []int64{120, 1, 1, 1, 1, 1}
	runSample(t, edges, cost, expect)
}

func TestSample2(t *testing.T) {
	edges := [][]int{{0, 1}, {0, 2}, {1, 3}, {1, 4}, {1, 5}, {2, 6}, {2, 7}, {2, 8}}
	cost := []int{1, 4, 2, 3, 5, 7, 8, -4, 2}
	expect := []int64{280, 140, 32, 1, 1, 1, 1, 1, 1}
	runSample(t, edges, cost, expect)
}
