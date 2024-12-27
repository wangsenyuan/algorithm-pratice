package main

import (
	"math"
	"testing"
)

func runSample(t *testing.T, a []float64) {
	b := solve(a)

	var sum int
	for i, x := range b {
		sum += x
		if math.Abs(a[i]-float64(x)) >= 1 {
			t.Fatalf("Sample result %v, not correct at %d, (%f)", b, i, a[i])
		}
	}

	if sum != 0 {
		t.Fatalf("Sample result %v, not sum to 0", b)
	}
}

func TestSample1(t *testing.T) {
	a := []float64{4.58413, 1.22491, -2.10517, -3.70387}
	runSample(t, a)
}

func TestSample2(t *testing.T) {
	a := []float64{-6.32509, 3.30066, -0.93878, 2.00000, 1.96321}
	runSample(t, a)
}

func TestSample3(t *testing.T) {
	a := []float64{-5.00000, -9.00000, 9.00000, 1.99728, 2.03313, 0.96959}
	runSample(t, a)
}
