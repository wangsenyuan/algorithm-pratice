package p1897

import "testing"

func runSample(t *testing.T, trips [][]int, target []int, expect bool) {
	res := mergeTriplets(trips, target)

	if res != expect {
		t.Errorf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	trips := [][]int{
		{2, 5, 3}, {1, 8, 4}, {1, 7, 5},
	}
	target := []int{2, 7, 5}
	runSample(t, trips, target, true)
}

func TestSample2(t *testing.T) {
	trips := [][]int{
		{1, 3, 4}, {2, 5, 8},
	}
	target := []int{2, 5, 8}
	runSample(t, trips, target, true)
}

func TestSample3(t *testing.T) {
	trips := [][]int{
		{1, 5, 7}, {3, 10, 1},
	}
	target := []int{3, 5, 7}
	runSample(t, trips, target, false)
}
