package main

import (
	"math"
	"testing"
)

func runSample(t *testing.T, R, K int, expect bool) {
	a, b := solve(R, K)
	if expect && a < 0 {
		t.Errorf("Sample %d %d, expect a result but got none", R, K)
		return
	}
	if expect {
		A, B := int64(a), int64(b)
		L := math.Sqrt(float64(A*A) + float64(B*B))
		if float64(R) < L {
			t.Errorf("Sample %d %d, result %d %d, out of cycle", R, K, a, b)
			return
		}
		if float64(R)-L > float64(K)/100 {
			t.Errorf("Sample %d %d, result %d %d, out of cycle not close enough", R, K, a, b)
		}
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 10, 16, true)
}

func TestSample2(t *testing.T) {
	runSample(t, 12, 1, false)
}
