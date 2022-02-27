package p2188

import "testing"

func runSample(t *testing.T, tiers [][]int, changeTime int, numLaps int, expect int) {
	res := minimumFinishTime(tiers, changeTime, numLaps)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	tires := [][]int{
		{2, 3}, {3, 4},
	}
	changeTime := 5
	numLaps := 4
	expect := 21
	runSample(t, tires, changeTime, numLaps, expect)
}

func TestSample2(t *testing.T) {
	tires := [][]int{
		{1, 10}, {2, 2}, {3, 4},
	}
	changeTime := 6
	numLaps := 5
	expect := 25
	runSample(t, tires, changeTime, numLaps, expect)
}

func TestSample3(t *testing.T) {
	tires := [][]int{
		{1, 10}, {2, 2}, {3, 4},
	}
	changeTime := 6
	numLaps := 1000
	expect := 5994
	runSample(t, tires, changeTime, numLaps, expect)
}
