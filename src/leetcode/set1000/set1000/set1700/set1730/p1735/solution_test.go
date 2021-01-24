package p1735

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, queries [][]int, expect []int) {
	res := waysToFillArray(queries)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	queries := [][]int{{2, 6}, {5, 1}, {73, 660}}
	expect := []int{4, 1, 50734910}
	runSample(t, queries, expect)
}

func TestSample2(t *testing.T) {
	queries := [][]int{{1, 1}, {2, 2}, {3, 3}, {4, 4}, {5, 5}}
	expect := []int{1, 2, 3, 10, 5}
	runSample(t, queries, expect)
}

func TestSample3(t *testing.T) {
	queries := [][]int{{4, 4}}
	expect := []int{10}
	runSample(t, queries, expect)
}
