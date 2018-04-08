package p815

import "testing"

func runSample(t *testing.T, routes [][]int, S int, T int, expect int) {
	res := numBusesToDestination(routes, S, T)
	if res != expect {
		t.Errorf("sample %v %d %d, expect %d, but got %d", routes, S, T, expect, res)
	}
}

func TestSample1(t *testing.T) {
	routes := [][]int{
		{1, 2, 7}, {3, 6, 7},
	}
	S := 1
	T := 6
	expect := 2
	runSample(t, routes, S, T, expect)
}

func TestSample2(t *testing.T) {
	routes := [][]int{{24}, {3, 6, 11, 14, 22}, {1, 23, 24}, {0, 6, 14}, {1, 3, 8, 11, 20}}
	S := 20
	T := 8
	expect := 1
	runSample(t, routes, S, T, expect)
}

func TestSample3(t *testing.T) {
	routes := [][]int{{0, 1, 6, 16, 22, 23}, {14, 15, 24, 32}, {4, 10, 12, 20, 24, 28, 33}, {1, 10, 11, 19, 27, 33}, {11, 23, 25, 28}, {15, 20, 21, 23, 29}, {29}}
	S := 4
	T := 21
	expect := 2
	runSample(t, routes, S, T, expect)
}

func TestSample4(t *testing.T) {
	routes := [][]int{{7, 12}, {4, 5, 15}, {6}, {15, 19}, {9, 12, 13}}
	S, T := 15, 12
	expect := -1
	runSample(t, routes, S, T, expect)
}
