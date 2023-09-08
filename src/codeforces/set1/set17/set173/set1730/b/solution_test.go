package main

import (
	"math"
	"testing"
)

func runSample(t *testing.T, x []int, y []int, expect float64) {
	res := solve(x, y)

	a := getValue(x, y, expect)
	b := getValue(x, y, res)

	if math.Abs(a-b) > 1e-6 {
		t.Fatalf("Sample expect %f, but got %f", expect, res)
	}
}

func getValue(x []int, y []int, x0 float64) float64 {
	var res float64 = -(1 << 20)

	for i := 0; i < len(x); i++ {
		cur := float64(x[i])
		if cur > x0 || math.Abs(cur-x0) < 1e-7 {
			res = math.Max(res, float64(y[i])+cur-x0)
		} else {
			res = math.Max(res, float64(y[i])+x0-cur)
		}
	}
	return res
}

func TestSample1(t *testing.T) {
	x := []int{0}
	y := []int{3}
	var expect float64 = 0
	runSample(t, x, y, expect)
}

func TestSample2(t *testing.T) {
	x := []int{3, 1}
	y := []int{0, 0}
	var expect float64 = 2
	runSample(t, x, y, expect)
}

func TestSample3(t *testing.T) {
	x := []int{1, 4}
	y := []int{0, 0}
	var expect float64 = 2.5
	runSample(t, x, y, expect)
}

func TestSample4(t *testing.T) {
	x := []int{1, 2, 3}
	y := []int{0, 0, 0}
	var expect float64 = 2
	runSample(t, x, y, expect)
}

func TestSample5(t *testing.T) {
	x := []int{1, 2, 3}
	y := []int{4, 1, 2}
	var expect float64 = 1
	runSample(t, x, y, expect)
}

func TestSample6(t *testing.T) {
	x := []int{3, 3, 3}
	y := []int{5, 3, 3}
	var expect float64 = 3
	runSample(t, x, y, expect)
}

func TestSample7(t *testing.T) {
	x := []int{5, 4, 7, 2, 10, 4}
	y := []int{3, 2, 5, 1, 4, 6}
	var expect float64 = 6
	runSample(t, x, y, expect)
}
