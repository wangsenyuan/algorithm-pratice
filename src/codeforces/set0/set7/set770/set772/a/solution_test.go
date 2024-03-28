package main

import (
	"math"
	"testing"
)

func runSample(t *testing.T, a []int, b []int, p int, expect float64) {
	res := solve(a, b, p)

	if math.Abs(res-expect) > 1e-5 {
		t.Fatalf("Sample expect %.7f, but got %.7f", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{2, 2}
	b := []int{2, 1000}
	p := 1
	expect := 2.0
	runSample(t, a, b, p, expect)
}

func TestSample2(t *testing.T) {
	a := []int{1}
	b := []int{1}
	p := 100
	expect := -1.0
	runSample(t, a, b, p, expect)
}

func TestSample3(t *testing.T) {
	a := []int{3, 4, 5, 6}
	b := []int{5, 3, 2, 1}
	p := 5
	expect := 0.5
	runSample(t, a, b, p, expect)
}
