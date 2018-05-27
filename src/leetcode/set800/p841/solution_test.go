package p841

import "testing"

func runSample(t *testing.T, rooms [][]int, expect bool) {
	res := canVisitAllRooms(rooms)

	if res != expect {
		t.Errorf("Sample %v, expect %t, but got %t", rooms, expect, res)
	}
}

func TestSample1(t *testing.T) {
	rooms := [][]int{
		{1}, {2}, {3}, {},
	}
	runSample(t, rooms, true)
}

func TestSample2(t *testing.T) {
	rooms := [][]int{
		{1, 3}, {3, 0, 1}, {2}, {0},
	}
	runSample(t, rooms, false)
}
