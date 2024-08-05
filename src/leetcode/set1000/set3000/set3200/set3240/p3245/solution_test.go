package p3245

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, colors []int, queries [][]int, expect []int) {
	res := numberOfAlternatingGroups(colors, queries)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	colors := []int{0, 1, 1, 0, 1}
	queries := [][]int{
		{2, 1, 0}, {1, 4},
	}
	// 0 0 1 0 1
	expect := []int{2}
	runSample(t, colors, queries, expect)
}

func TestSample2(t *testing.T) {
	colors := []int{0, 0, 1, 0, 1, 1}
	queries := [][]int{
		{1, 3}, {2, 3, 0}, {1, 5},
	}
	expect := []int{2, 0}
	runSample(t, colors, queries, expect)
}

func TestSample3(t *testing.T) {
	colors := []int{0, 1, 0, 1}
	queries := [][]int{
		{1, 3}, {2, 2, 1}, {1, 3}, {1, 3},
	}
	// 0 1 1 1
	expect := []int{4, 1, 1}
	runSample(t, colors, queries, expect)
}

func TestSample4(t *testing.T) {
	colors := []int{0, 1, 1, 0}
	queries := [][]int{
		{2, 1, 1}, {1, 3}, {2, 3, 1}, {2, 0, 1}, {2, 3, 0},
	}
	// 0 1 1 1
	expect := []int{0}
	runSample(t, colors, queries, expect)
}

func TestSample5(t *testing.T) {
	colors := []int{1, 1, 0, 1}
	queries := [][]int{
		{2, 1, 0}, {2, 2, 1}, {1, 3}, {2, 2, 0}, {1, 3},
	}
	// 1,0,0,1
	// 1,0,1,1
	// 1,0,0,1
	expect := []int{1, 0}
	runSample(t, colors, queries, expect)
}
