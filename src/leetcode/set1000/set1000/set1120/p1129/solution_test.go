package p1129

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, red_edges [][]int, blue_edges [][]int, expect []int) {
	res := shortestAlternatingPaths(n, red_edges, blue_edges)
	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample %d %v %v, expect %v, but got %v", n, red_edges, blue_edges, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	red_edges := [][]int{
		{0, 1}, {1, 2},
	}
	blue_edges := [][]int{}
	expect := []int{0, 1, -1}

	runSample(t, n, red_edges, blue_edges, expect)
}

func TestSample2(t *testing.T) {
	n := 3
	red_edges := [][]int{
		{0, 1},
	}
	blue_edges := [][]int{{2, 1}}
	expect := []int{0, 1, -1}

	runSample(t, n, red_edges, blue_edges, expect)
}

func TestSample3(t *testing.T) {
	n := 3
	red_edges := [][]int{
		{1, 0},
	}
	blue_edges := [][]int{{2, 1}}
	expect := []int{0, -1, -1}

	runSample(t, n, red_edges, blue_edges, expect)
}

func TestSample4(t *testing.T) {
	n := 3
	red_edges := [][]int{
		{0, 1}, {0, 2},
	}
	blue_edges := [][]int{{1, 0}}
	expect := []int{0, 1, 1}

	runSample(t, n, red_edges, blue_edges, expect)
}
