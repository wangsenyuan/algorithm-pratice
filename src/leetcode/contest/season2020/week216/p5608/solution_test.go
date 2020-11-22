package p5608

import "testing"

func runSample(t *testing.T, tasks [][]int, expect int) {
	res := minimumEffort(tasks)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	tasks := [][]int{
		{1, 2}, {2, 4}, {4, 8},
	}
	expect := 8
	runSample(t, tasks, expect)
}

func TestSample2(t *testing.T) {
	tasks := [][]int{
		{1, 3}, {2, 4}, {10, 11}, {10, 12}, {8, 9},
	}
	expect := 32
	runSample(t, tasks, expect)
}

func TestSample3(t *testing.T) {
	tasks := [][]int{
		{1, 7}, {2, 8}, {3, 9}, {4, 10}, {5, 11}, {6, 12},
	}
	expect := 27
	runSample(t, tasks, expect)
}
