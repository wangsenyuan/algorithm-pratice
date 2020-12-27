package p1700

import (
	"math"
	"testing"
)

func runSample(t *testing.T, cust [][]int, expect float64) {
	res := averageWaitingTime(cust)

	if math.Abs(res-expect) > 1e-6 {
		t.Errorf("Sample expect %f, but got %f", expect, res)
	}
}

func TestSample1(t *testing.T) {
	arr := [][]int{
		{1, 2}, {2, 5}, {4, 3},
	}
	expect := 5.0
	runSample(t, arr, expect)
}

func TestSample2(t *testing.T) {
	arr := [][]int{
		{5, 2}, {5, 4}, {10, 3}, {20, 1},
	}
	expect := 3.25
	runSample(t, arr, expect)
}
