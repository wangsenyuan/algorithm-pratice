package p3241

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, edges [][]int, expect []int) {
	res := timeTaken(edges)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	edges := [][]int{
		{0, 1},
		{0, 2},
	}
	expect := []int{2, 4, 3}
	runSample(t, edges, expect)
}

func TestSample2(t *testing.T) {
	edges := [][]int{
		{0, 1},
	}
	expect := []int{1, 2}
	runSample(t, edges, expect)
}

func TestSample3(t *testing.T) {
	edges := [][]int{
		{2, 4}, {0, 1}, {2, 3}, {0, 2},
	}
	expect := []int{4, 6, 3, 5, 5}
	runSample(t, edges, expect)
}

func TestSample4(t *testing.T) {
	edges := [][]int{
		{5, 4}, {2, 1}, {4, 2}, {1, 0}, {3, 1},
	}
	expect := []int{6, 5, 3, 6, 5, 7}
	runSample(t, edges, expect)
}
