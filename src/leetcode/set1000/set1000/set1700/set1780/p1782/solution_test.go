package p1782

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, E [][]int, queries []int, expect []int) {
	res := countPairs(n, E, queries)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 4
	E := [][]int{
		{1, 2}, {2, 4}, {1, 3}, {2, 3}, {2, 1},
	}
	Q := []int{2, 3}
	expect := []int{6, 5}
	runSample(t, n, E, Q, expect)
}

func TestSample2(t *testing.T) {
	n := 5
	E := [][]int{
		{1, 5}, {1, 5}, {3, 4}, {2, 5}, {1, 3}, {5, 1}, {2, 3}, {2, 5},
	}
	Q := []int{1, 2, 3, 4, 5}
	expect := []int{10, 10, 9, 8, 6}
	runSample(t, n, E, Q, expect)
}

func TestSample3(t *testing.T) {
	n := 5
	E := [][]int{
		{4, 5}, {1, 3}, {1, 4},
	}
	Q := []int{0, 1, 0, 0, 1, 1, 0, 0, 0, 0, 1, 0, 0, 1, 0, 1, 1, 2}
	expect := []int{10, 8, 10, 10, 8, 8, 10, 10, 10, 10, 8, 10, 10, 8, 10, 8, 8, 3}
	runSample(t, n, E, Q, expect)
}

func TestSample4(t *testing.T) {
	n := 8
	E := [][]int{
		{2, 5}, {5, 4}, {4, 3}, {7, 3}, {5, 8}, {6, 1}, {7, 5}, {6, 1}, {7, 5}, {3, 1}, {6, 4}, {6, 8}, {6, 5}, {5, 6}, {8, 1}, {1, 6}, {1, 3}, {5, 6}, {8, 3}, {6, 2},
	}
	Q := []int{2, 5, 4, 3, 0, 4, 4, 3, 3, 2, 3, 1, 5, 7, 5, 1, 6, 18, 14}
	expect := []int{28, 26, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 26, 19, 26, 28, 24, 0, 0}
	runSample(t, n, E, Q, expect)
}
