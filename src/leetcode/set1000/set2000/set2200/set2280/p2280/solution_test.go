package p2280

import (
	"testing"
)

func runSample(t *testing.T, prices [][]int, expect int) {
	res := minimumLines(prices)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	prices := [][]int{
		{1, 7}, {2, 6}, {3, 5}, {4, 4}, {5, 4}, {6, 3}, {7, 2}, {8, 1},
	}
	expect := 3
	runSample(t, prices, expect)
}

func TestSample2(t *testing.T) {
	prices := [][]int{
		{3, 4}, {1, 2}, {7, 8}, {2, 3},
	}
	expect := 1
	runSample(t, prices, expect)
}
