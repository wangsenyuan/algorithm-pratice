package e

import "testing"

func runSample(t *testing.T, tasks [][]int, expect int) {
	res := processTasks(tasks)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	tasks := [][]int{
		{1, 3, 2}, {2, 5, 3}, {5, 6, 2},
	}
	expect := 4
	runSample(t, tasks, expect)
}

func TestSample2(t *testing.T) {
	tasks := [][]int{
		{2, 3, 1}, {5, 5, 1}, {5, 6, 2},
	}
	expect := 3
	runSample(t, tasks, expect)
}

func TestSample3(t *testing.T) {
	tasks := [][]int{
		{0, 0, 1}, {1, 1, 1}, {2, 2, 1},
	}
	expect := 3
	runSample(t, tasks, expect)
}

func TestSample4(t *testing.T) {
	tasks := [][]int{
		{29, 70, 3}, {43, 48, 3}, {30, 83, 27}, {6, 75, 26}, {25, 96, 41}, {2, 80, 44}, {74, 88, 15}, {52, 97, 17}, {33, 50, 2}, {66, 82, 8},
	}
	expect := 52
	runSample(t, tasks, expect)
}

func TestSample5(t *testing.T) {
	tasks := [][]int{
		{10, 42, 6}, {47, 81, 35}, {38, 58, 13}, {39, 48, 4}, {12, 62, 24}, {54, 73, 1}, {59, 96, 34}, {65, 91, 20}}
	expect := 54
	runSample(t, tasks, expect)
}
