package p857

import (
	"math"
	"testing"
)

func runSample(t *testing.T, quality []int, wage []int, K int, expect float64) {
	res := mincostToHireWorkers(quality, wage, K)

	if math.Abs(res-expect) > 1e-5 {
		t.Errorf("Sample %v %v %d, expect %f, but got %f", quality, wage, K, expect, res)
	}
}

func TestSample1(t *testing.T) {
	quality := []int{10, 20, 5}
	wage := []int{70, 50, 30}
	K := 2
	runSample(t, quality, wage, K, 105)
}

func TestSample2(t *testing.T) {
	quality := []int{3, 1, 10, 10, 1}
	wage := []int{4, 8, 2, 2, 7}
	K := 3
	runSample(t, quality, wage, K, 30.66667)
}
