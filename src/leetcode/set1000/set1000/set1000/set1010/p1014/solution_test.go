package p1014

import "testing"

func runSample(t *testing.T, weights []int, D int, expect int) {
	res := shipWithinDays(weights, D)
	if res != expect {
		t.Errorf("Sample %v %d, expect %d, but got %d", weights, D, expect, res)
	}
}

func TestSample1(t *testing.T) {
	weights := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	D := 5
	expect := 15
	runSample(t, weights, D, expect)
}

func TestSample2(t *testing.T) {
	weights := []int{3, 2, 2, 4, 1, 4}
	D := 3
	expect := 6
	runSample(t, weights, D, expect)
}

func TestSample3(t *testing.T) {
	weights := []int{1, 2, 3, 1, 1}
	D := 4
	expect := 3
	runSample(t, weights, D, expect)
}
