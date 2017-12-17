package p746

import "testing"

func runSample(t *testing.T, cost []int, expect int) {
	res := minCostClimbingStairs(cost)
	if res != expect {
		t.Errorf("climbing stairs on cost %v, should give %d, but got %d", cost, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, []int{10, 15, 20}, 15)
}

func TestSample2(t *testing.T) {
	runSample(t, []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}, 6)
}
