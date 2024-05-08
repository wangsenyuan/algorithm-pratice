package main

import (
	"math"
	"testing"
)

func runSample(t *testing.T, x []int, v []int, T int, expect float64) {
	res := solve(x, v, T)

	if math.Abs(res-expect)/math.Max(expect, 1) > eps {
		t.Fatalf("Sample expect %f, but got %f", expect, res)
	}
}

func TestSample1(t *testing.T) {
	v := []int{3, 10}
	x := []int{50, 150}
	T := 100
	var expect float64 = 6
	runSample(t, x, v, T, expect)
}

func TestSample2(t *testing.T) {
	v := []int{5, 5, 30}
	x := []int{6, 6, 10}
	T := 9
	var expect float64 = 40
	runSample(t, x, v, T, expect)
}

func TestSample3(t *testing.T) {
	v := []int{1, 3}
	x := []int{10, 15}
	T := 12
	var expect float64 = 1.666666666666667
	runSample(t, x, v, T, expect)
}

func TestSample4(t *testing.T) {
	v := []int{70, 97, 14, 31, 83, 22, 83, 56, 19, 87, 59, 7, 7, 89, 24, 82, 34, 40, 6, 24}
	x := []int{10, 4, 47, 46, 11, 18, 32, 55, 16, 32, 53, 37, 43, 32, 41, 46, 57, 14, 60, 44}
	T := 30
	var expect float64 = 916.518518518518519
	runSample(t, x, v, T, expect)
}
