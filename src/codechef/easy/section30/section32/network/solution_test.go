package main

import (
	"math"
	"testing"
)

func runSample(t *testing.T, access_points [][]int, computers [][]int, expect []float64) {
	res := solve(access_points, computers)

	for i := 0; i < len(res); i++ {
		if math.Abs(res[i]-expect[i]) > 1e-3 {
			t.Fatalf("Sample expect %v, but got %v", expect, res)
		}
	}
}

func TestSample1(t *testing.T) {
	access_points := [][]int{
		{5, 5, 2},
		{6, 9, 2},
	}
	computers := [][]int{
		{3, 5},
		{5, 2},
		{4, 2},
	}
	expect := []float64{2, 3, 5}
	runSample(t, access_points, computers, expect)
}

func TestSample2(t *testing.T) {
	access_points := [][]int{
		{3, 6, 2},
		{10, 9, 2},
	}
	computers := [][]int{
		{2, 5},
		{3, 8},
		{4, 3},
		{6, 4},
	}
	expect := []float64{1.414, 2.000, 7.071, 7.071}
	runSample(t, access_points, computers, expect)
}

func TestSample3(t *testing.T) {
	access_points := [][]int{
		{10, 5, 1},
		{6, 8, 1},
	}
	computers := [][]int{
		{3, 3},
	}
	expect := []float64{5.831}
	runSample(t, access_points, computers, expect)
}
