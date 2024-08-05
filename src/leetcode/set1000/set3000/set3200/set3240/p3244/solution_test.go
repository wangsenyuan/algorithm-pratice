package p3244

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, queries [][]int, expect []int) {
	res := shortestDistanceAfterQueries(n, queries)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 5
	queries := [][]int{
		{2, 4}, {0, 2}, {0, 4},
	}
	expect := []int{3, 2, 1}
	runSample(t, n, queries, expect)
}

func TestSample2(t *testing.T) {
	n := 4
	queries := [][]int{
		{0, 3}, {0, 2},
	}
	expect := []int{1, 1}
	runSample(t, n, queries, expect)
}

func TestSample3(t *testing.T) {
	n := 4
	queries := [][]int{
		{0, 3}, {0, 2},
	}
	expect := []int{1, 1}
	runSample(t, n, queries, expect)
}

func TestSample4(t *testing.T) {
	n := 5
	queries := [][]int{
		{1, 3}, {2, 4},
	}
	expect := []int{3, 3}
	runSample(t, n, queries, expect)
}

func TestSample5(t *testing.T) {
	n := 5
	queries := [][]int{
		{0, 2}, {0, 4},
	}
	expect := []int{3, 1}
	runSample(t, n, queries, expect)
}
