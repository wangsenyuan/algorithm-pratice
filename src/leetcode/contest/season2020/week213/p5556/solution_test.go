package p5556

import "testing"

func runSample(t *testing.T, heights []int, bricks int, ladders int, expect int) {
	res := furthestBuilding(heights, bricks, ladders)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	heights := []int{4, 2, 7, 6, 9, 14, 12}
	bricks := 5
	ladders := 1
	expect := 4
	runSample(t, heights, bricks, ladders, expect)
}

func TestSample2(t *testing.T) {
	heights := []int{4, 12, 2, 7, 3, 18, 20, 3, 19}
	bricks := 10
	ladders := 2
	expect := 7
	runSample(t, heights, bricks, ladders, expect)
}

func TestSample3(t *testing.T) {
	heights := []int{14, 3, 19, 3}
	bricks := 17
	ladders := 0
	expect := 3
	runSample(t, heights, bricks, ladders, expect)
}
