package p1833

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, tasks [][]int, expect []int) {
	res := getOrder(tasks)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	tasks := [][]int{
		{1, 2}, {2, 4}, {3, 2}, {4, 1},
	}
	expect := []int{
		0, 2, 3, 1,
	}
	runSample(t, tasks, expect)
}

func TestSample2(t *testing.T) {
	tasks := [][]int{
		{7, 10}, {7, 12}, {7, 5}, {7, 4}, {7, 2},
	}
	expect := []int{
		4, 3, 2, 0, 1,
	}
	runSample(t, tasks, expect)
}

func TestSample3(t *testing.T) {
	tasks := [][]int{
		{19, 13}, {16, 9}, {21, 10}, {32, 25}, {37, 4}, {49, 24}, {2, 15}, {38, 41}, {37, 34}, {33, 6}, {45, 4}, {18, 18}, {46, 39}, {12, 24},
	}
	expect := []int{
		6, 1, 2, 9, 4, 10, 0, 11, 5, 13, 3, 8, 12, 7,
	}
	runSample(t, tasks, expect)
}

func TestSample4(t *testing.T) {
	tasks := [][]int{
		{35, 36}, {11, 7}, {15, 47}, {34, 2}, {47, 19}, {16, 14}, {19, 8}, {7, 34}, {38, 15}, {16, 18}, {27, 22}, {7, 15}, {43, 2}, {10, 5}, {5, 4}, {3, 11},
	}
	expect := []int{
		15, 14, 13, 1, 6, 3, 5, 12, 8, 11, 9, 4, 10, 7, 0, 2,
	}
	runSample(t, tasks, expect)
}
