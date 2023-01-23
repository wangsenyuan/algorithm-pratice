package main

import (
	"math"
	"strconv"
	"testing"
)

func runSample(t *testing.T, P, V []int, k int, expect float64) {
	res := solve(P, V, k)

	rf, _ := strconv.ParseFloat(res, 64)

	if math.Abs(rf-expect) > 1e-9 {
		t.Errorf("Sample expect %.10f, but got %.10f", expect, rf)
	}
}

func TestSample1(t *testing.T) {
	P := []int{1, 10}
	V := []int{1, 10}
	k := 2
	expect := 0.32
	runSample(t, P, V, k, expect)
}

func TestSample2(t *testing.T) {
	P := []int{5, 6}
	V := []int{8, 10}
	k := 1
	expect := 1.0
	runSample(t, P, V, k, expect)
}

func TestSample3(t *testing.T) {
	P := []int{1, 1000000000}
	V := []int{1, 1000000000}
	k := 47
	expect := 0.000000010664
	runSample(t, P, V, k, expect)
}
