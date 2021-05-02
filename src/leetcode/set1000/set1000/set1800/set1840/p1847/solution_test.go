package p1847

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, rooms [][]int, queries [][]int, expect []int) {
	res := closestRoom(rooms, queries)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	rooms := [][]int{
		{2, 2}, {1, 2}, {3, 2},
	}
	queries := [][]int{
		{3, 1}, {3, 3}, {5, 2},
	}
	expect := []int{3, -1, 3}
	runSample(t, rooms, queries, expect)
}

func TestSample2(t *testing.T) {
	rooms := [][]int{
		{1, 4}, {2, 3}, {3, 5}, {4, 1}, {5, 2},
	}
	queries := [][]int{
		{2, 3}, {2, 4}, {2, 5},
	}
	expect := []int{2, 1, 3}
	runSample(t, rooms, queries, expect)
}
