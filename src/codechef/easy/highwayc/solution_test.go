package main

import (
	"math"
	"testing"
)

func runSample(t *testing.T, N int, S int, Y int, V []int, D []int, P []int, C []int, expect float64) {
	res := solve(N, S, Y, V, D, P, C)
	if math.Abs(res-expect) > 1e-6 {
		t.Errorf("sample %d %d %d %v %v %v %v, expect %f, but got %f", N, S, Y, V, D, P, C, expect, res)
	}
}

func TestSample1(t *testing.T) {
	N, S, Y := 2, 6, 6
	V := []int{5, 10}
	D := []int{1, 0}
	P := []int{-1, 23}
	C := []int{4, 7}
	expect := float64(4.0)
	runSample(t, N, S, Y, V, D, P, C, expect)
}

func TestSample2(t *testing.T) {
	N, S, Y := 1, 8, 177
	V := []int{190}
	D := []int{0}
	P := []int{4201}
	C := []int{21}
	expect := float64(44.346053)
	runSample(t, N, S, Y, V, D, P, C, expect)
}

func TestSample3(t *testing.T) {
	N, S, Y := 1, 1, 1000
	V := []int{100}
	D := []int{1}
	P := []int{-100000}
	C := []int{10}
	expect := float64(2000.100000)
	runSample(t, N, S, Y, V, D, P, C, expect)
}
