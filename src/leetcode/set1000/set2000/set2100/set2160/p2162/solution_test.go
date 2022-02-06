package p2162

import (
	"testing"
)

func runSample(t *testing.T, startAt int, moveCost int, pushCost int, targetSeconds int, expect int) {
	res := minCostSetTime(startAt, moveCost, pushCost, targetSeconds)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	startAt := 1
	moveCost := 2
	pushCost := 1
	targetSeconds := 600
	expect := 6
	runSample(t, startAt, moveCost, pushCost, targetSeconds, expect)
}

func TestSample2(t *testing.T) {
	startAt := 0
	moveCost := 1
	pushCost := 2
	targetSeconds := 76
	expect := 6
	runSample(t, startAt, moveCost, pushCost, targetSeconds, expect)
}
