package p1176

import "testing"

func runSample(t *testing.T, calories []int, k int, lower int, upper int, expect int) {
	res := dietPlanPerformance(calories, k, lower, upper)

	if res != expect {
		t.Errorf("Sample %v %d %d %d, expect %d, but got %d", calories, k, lower, upper, expect, res)
	}
}

func TestSample1(t *testing.T) {
	calories := []int{3, 2}
	runSample(t, calories, 2, 0, 1, 1)
}
