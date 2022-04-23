package main

import (
	"math"
	"testing"
)

func runSample(t *testing.T, n int, B []int, Q [][]int, expect []float64) {
	res := solve(n, B, Q)

	for i := 0; i < len(res); i++ {
		if math.Abs(res[i]-expect[i]) > 1e-7 {
			t.Errorf("Sample result %v, failed at %d, expect %f", res, i, expect[i])
		}
	}
}

func TestSample1(t *testing.T) {
	n := 1
	B := []int{5}
	Q := [][]int{{0, 0}}
	expect := []float64{5}
	runSample(t, n, B, Q, expect)
}

func TestSample2(t *testing.T) {
	n := 2
	B := []int{3, 5}
	Q := [][]int{{0, 1}}
	expect := []float64{4}
	runSample(t, n, B, Q, expect)
}

func TestSample3(t *testing.T) {
	n := 18
	B := []int{3, 4, 2, 1, 5, 7, 9, 7, 10, 5, 12, 3, 1, 1, 2, 1, 3, 2}
	Q := [][]int{{4, 10}}
	expect := []float64{9.0}
	runSample(t, n, B, Q, expect)
}

func TestSample4(t *testing.T) {
	n := 6
	B := []int{3, 4, 5, 1, 2, 3}
	Q := [][]int{{0, 2}}
	expect := []float64{6}
	runSample(t, n, B, Q, expect)
}
