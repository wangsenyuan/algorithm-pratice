package p2744

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, logs [][]int, x int, queries []int, expect []int) {
	res := countServers(n, logs, x, queries)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	logs := [][]int{
		{1, 3}, {2, 6}, {1, 5},
	}
	x := 5
	queries := []int{10, 11}
	expect := []int{1, 2}
	runSample(t, n, logs, x, queries, expect)
}

func TestSample2(t *testing.T) {
	n := 3
	logs := [][]int{
		{2, 4}, {2, 1}, {1, 2}, {3, 1},
	}
	x := 2
	queries := []int{3, 4}
	expect := []int{0, 1}
	runSample(t, n, logs, x, queries, expect)
}
