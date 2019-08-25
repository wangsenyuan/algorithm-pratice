package p1010

import "testing"

func runSample(t *testing.T, costs [][]int, expect int) {
	res := twoCitySchedCost(costs)

	if res != expect {
		t.Errorf("Sample %v, expect %d, but got %d", costs, expect, res)
	}
}

func TestSample1(t *testing.T) {
	costs := [][]int{
		{10, 20}, {30, 200}, {400, 50}, {30, 20},
	}
	expect := 110
	runSample(t, costs, expect)
}
