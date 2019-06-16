package p813

import (
	"testing"
	"math"
)

func runSample(t *testing.T, A []int, K int, expect float64) {
	res := largestSumOfAverages(A, K)

	if math.Abs(expect-res) > 1e-7 {
		t.Errorf("Sample %v %d, expect %f, but got %f", A, K, expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{9, 1, 2, 3, 9}
	K := 3
	runSample(t, A, K, 20)
}
