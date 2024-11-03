package p3342

import "testing"

func runSample(t *testing.T, moveTime [][]int, expect int) {
	res := minTimeToReach(moveTime)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	moveTime := [][]int{
		{0, 4}, {4, 4},
	}
	expect := 7
	runSample(t, moveTime, expect)
}

func TestSample2(t *testing.T) {
	moveTime := [][]int{
		{0, 0, 0, 0}, {0, 0, 0, 0},
	}
	expect := 6
	runSample(t, moveTime, expect)
}

func TestSample3(t *testing.T) {
	moveTime := [][]int{
		{0, 1}, {1, 2},
	}
	expect := 4
	runSample(t, moveTime, expect)
}
