package p2462

import "testing"

func runSample(t *testing.T, costs []int, k int, candidate int, expect int64) {
	res := totalCost(costs, k, candidate)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	costs := []int{17, 12, 10, 2, 7, 2, 11, 20, 8}
	k := 3
	candidates := 4
	var expect int64 = 11
	runSample(t, costs, k, candidates, expect)
}

func TestSample2(t *testing.T) {
	costs := []int{1, 2, 4, 1}
	k := 3
	candidates := 3
	var expect int64 = 4
	runSample(t, costs, k, candidates, expect)
}

func TestSample3(t *testing.T) {
	costs := []int{31, 25, 72, 79, 74, 65, 84, 91, 18, 59, 27, 9, 81, 33, 17, 58}
	k := 11
	candidates := 2
	var expect int64 = 423
	runSample(t, costs, k, candidates, expect)
}
