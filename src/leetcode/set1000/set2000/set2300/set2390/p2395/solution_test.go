package p2395

import "testing"

func runSample(t *testing.T, charge []int, cost []int, budget int64, expect int) {
	res := maximumRobots(charge, cost, budget)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	chargeTimes := []int{3, 6, 1, 3, 4}
	runningCosts := []int{2, 1, 3, 4, 5}
	budget := 25
	expect := 3
	runSample(t, chargeTimes, runningCosts, int64(budget), expect)
}

func TestSample2(t *testing.T) {
	chargeTimes := []int{11, 12, 19}
	runningCosts := []int{10, 8, 7}
	budget := 9
	expect := 0
	runSample(t, chargeTimes, runningCosts, int64(budget), expect)
}

func TestSample3(t *testing.T) {
	chargeTimes := []int{11, 12, 74, 67, 37, 87, 42, 34, 18, 90, 36, 28, 34, 20}
	runningCosts := []int{18, 98, 2, 84, 7, 57, 54, 65, 59, 91, 7, 23, 94, 20}
	budget := 937
	expect := 4
	runSample(t, chargeTimes, runningCosts, int64(budget), expect)
}

func TestSample4(t *testing.T) {
	chargeTimes := []int{19, 63, 21, 8, 5, 46, 56, 45, 54, 30, 92, 63, 31, 71, 87, 94, 67, 8, 19, 89, 79, 25}
	runningCosts := []int{91, 92, 39, 89, 62, 81, 33, 99, 28, 99, 86, 19, 5, 6, 19, 94, 65, 86, 17, 10, 8, 42}
	budget := 85
	expect := 1
	runSample(t, chargeTimes, runningCosts, int64(budget), expect)
}
