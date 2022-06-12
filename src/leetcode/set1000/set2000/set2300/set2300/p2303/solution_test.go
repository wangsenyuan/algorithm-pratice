package p2303

import (
	"math"
	"testing"
)

func runSample(t *testing.T, brackets [][]int, income int, expect float64) {
	res := calculateTax(brackets, income)
	if math.Abs(res-expect) > 1e-6 {
		t.Errorf("Sample expect %f, but got %f", expect, res)
	}
}

func TestSample1(t *testing.T) {
	brackets := [][]int{{3, 50}, {7, 10}, {12, 25}}
	income := 10
	expect := 2.65
	runSample(t, brackets, income, expect)
}

func TestSample2(t *testing.T) {
	brackets := [][]int{{1, 0}, {4, 25}, {5, 50}}
	income := 2
	expect := 0.25
	runSample(t, brackets, income, expect)
}
