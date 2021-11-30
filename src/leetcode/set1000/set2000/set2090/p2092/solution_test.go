package p2092

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, meetings [][]int, first int, expect []int) {
	res := findAllPeople(n, meetings, first)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 6
	meetings := [][]int{
		{1, 2, 5}, {2, 3, 8}, {1, 5, 10},
	}
	first := 1
	expect := []int{0, 1, 2, 3, 5}
	runSample(t, n, meetings, first, expect)
}

func TestSample2(t *testing.T) {
	n := 4
	meetings := [][]int{
		{3, 1, 3}, {1, 2, 2}, {0, 3, 3},
	}
	first := 3
	expect := []int{0, 1, 3}
	runSample(t, n, meetings, first, expect)
}

func TestSample4(t *testing.T) {
	n := 5
	meetings := [][]int{
		{3, 4, 2}, {1, 2, 1}, {2, 3, 1},
	}
	first := 1
	expect := []int{0, 1, 2, 3, 4}
	runSample(t, n, meetings, first, expect)
}

func TestSample5(t *testing.T) {
	n := 6
	meetings := [][]int{
		{0, 2, 1}, {1, 3, 1}, {4, 5, 1},
	}
	first := 1
	expect := []int{0, 1, 2, 3}
	runSample(t, n, meetings, first, expect)
}
