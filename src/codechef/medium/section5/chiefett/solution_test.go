package main

import (
	"math"
	"testing"
)

func runSample(t *testing.T, n, k int, prices []int, discount []int, expect float64) {
	res := solve(n, k, prices, discount)

	if math.Abs(res-expect) > 1e-4 {
		t.Errorf("Sample %d %d %v %v, expect %f, but got %f", n, k, prices, discount, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, k := 3, 2
	prices := []int{100, 200, 300}
	discount := []int{10, 20}
	expect := 66.6667
	runSample(t, n, k, prices, discount, expect)
}

func TestSample2(t *testing.T) {
	n, k := 4, 3
	prices := []int{100, 200, 300, 400}
	discount := []int{10, 20, 30}
	expect := 175.000
	runSample(t, n, k, prices, discount, expect)
}
