package p2463

import "testing"

func runSample(t *testing.T, robot []int, factory [][]int, expect int) {
	res := int(minimumTotalDistance(robot, factory))

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	robot := []int{0, 4, 6}
	factory := [][]int{
		{2, 2}, {6, 2},
	}
	expect := 4
	runSample(t, robot, factory, expect)
}
func TestSample2(t *testing.T) {
	robot := []int{1, -1}
	factory := [][]int{
		{-2, 1}, {2, 1},
	}
	expect := 2
	runSample(t, robot, factory, expect)
}

func TestSample3(t *testing.T) {
	robot := []int{9, 11, 99, 101}
	factory := [][]int{
		{10, 1}, {7, 1}, {14, 1}, {100, 1}, {96, 1}, {103, 1},
	}
	expect := 6
	runSample(t, robot, factory, expect)
}
