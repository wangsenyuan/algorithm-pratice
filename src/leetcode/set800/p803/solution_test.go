package p803

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, grid [][]int, hits [][]int, expect []int) {
	res := hitBricks(grid, hits)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("sample %v %v, expect %v, but got %v", grid, hits, expect, res)
	}
}

func TestSample1(t *testing.T) {
	grid := [][]int{{1, 0, 0, 0}, {1, 1, 1, 0}}
	hits := [][]int{{1, 0}}
	expect := []int{2}
	runSample(t, grid, hits, expect)
}

func TestSample2(t *testing.T) {
	grid := [][]int{{1, 0, 0, 0}, {1, 1, 0, 0}}
	hits := [][]int{{1, 1}, {1, 0}}
	expect := []int{0, 0}
	runSample(t, grid, hits, expect)
}

func TestSample3(t *testing.T) {
	grid := [][]int{{1, 1, 1, 0, 1, 1, 1, 1}, {1, 0, 0, 0, 0, 1, 1, 1}, {1, 1, 1, 0, 0, 0, 1, 1}, {1, 1, 0, 0, 0, 0, 0, 0}, {1, 0, 0, 0, 0, 0, 0, 0}, {1, 0, 0, 0, 0, 0, 0, 0}}
	hits := [][]int{{4, 6}, {3, 0}, {2, 3}, {2, 6}, {4, 1}, {5, 2}, {2, 1}}
	expect := []int{0, 2, 0, 0, 0, 0, 2}
	runSample(t, grid, hits, expect)
}

func TestSample4(t *testing.T) {
	grid := [][]int{{0, 0, 0, 0, 0}, {1, 1, 1, 1, 1}, {1, 1, 1, 1, 1}, {0, 0, 0, 0, 0}}
	hits := [][]int{{1, 1}, {1, 2}}
	expect := []int{9, 0}
	runSample(t, grid, hits, expect)
}

func TestSample5(t *testing.T) {
	graph := [][]int{
		{0, 1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1, 1},
		{0, 0, 1, 0, 0, 0},
		{0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0},
	}
	hits := [][]int{{1, 3}, {3, 5}, {0, 3}, {3, 3}, {1, 1}, {4, 2}, {1, 0}, {3, 0}, {4, 5}, {2, 1}, {4, 4}, {4, 0}, {2, 4}, {2, 5}, {3, 4}, {0, 5}, {0, 4}, {3, 2}, {1, 5}, {4, 1}, {2, 2}, {0, 2}}

	expect := []int{0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 2, 0, 0, 0, 0, 1}
	runSample(t, graph, hits, expect)
}
