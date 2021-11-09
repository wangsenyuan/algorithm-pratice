package p2065

import "testing"

func runSample(t *testing.T, values []int, edges [][]int, maxTime int, expect int) {
	res := maximalPathQuality(values, edges, maxTime)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	values := []int{0, 32, 10, 43}
	edges := [][]int{
		{0, 1, 10}, {1, 2, 15}, {0, 3, 10},
	}
	maxTime := 49
	expect := 75
	runSample(t, values, edges, maxTime, expect)
}
