package main

import (
	"math"
	"testing"
)

func runSample(t *testing.T, n int, m int, a []int, expect float64) {
	res := solve(n, m, a)

	if math.Abs(res-expect) > 1e-7 {
		t.Fatalf("Sample expect %.8f, but got %.8f", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, m := 2, 1
	a := []int{2}
	expect := 1.0
	runSample(t, n, m, a, expect)
}

func TestSample2(t *testing.T) {
	n, m := 2, 2
	a := []int{1, 1}
	expect := 1.5
	runSample(t, n, m, a, expect)
}

func TestSample3(t *testing.T) {
	n, m := 2, 3
	a := []int{1, 1, 1}
	expect := 1.33333333333333350000
	runSample(t, n, m, a, expect)
}

func TestSample4(t *testing.T) {
	n, m := 7, 5
	a := []int{1, 1, 2, 3, 1}
	expect := 2.50216960000000070000
	runSample(t, n, m, a, expect)
}

func TestSample5(t *testing.T) {
	n, m := 10, 4
	a := []int{8, 4, 7, 6}
	expect := 1.08210754
	runSample(t, n, m, a, expect)
}
