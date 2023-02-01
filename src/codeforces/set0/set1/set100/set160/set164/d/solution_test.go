package main

import (
	"testing"
)

func runSample(t *testing.T, k int, points [][]int, expect []int) {
	res := solve(k, points)
	x := minimumDiameter(points, expect)
	y := minimumDiameter(points, res)
	if x != y {
		t.Errorf("Sample expect %d, but got %d", x, y)
	}
}

func minimumDiameter(points [][]int, rem []int) int64 {
	m := make(map[int]bool)

	for _, num := range rem {
		m[num-1] = true
	}

	var res int64

	for i := 0; i < len(points); i++ {
		if !m[i] {
			for j := i + 1; j < len(points); j++ {
				if !m[j] {
					res = max(res, distance(points[i], points[j]))
				}
			}
		}
	}

	return res
}

func TestSample1(t *testing.T) {
	points := [][]int{
		{1, 2},
		{0, 0},
		{2, 2},
		{1, 1},
		{3, 3},
	}
	k := 2
	expect := []int{5, 2}
	runSample(t, k, points, expect)
}

func TestSample2(t *testing.T) {
	points := [][]int{
		{0, 0},
		{0, 0},
		{1, 1},
		{1, 1},
	}
	k := 1
	expect := []int{3}
	runSample(t, k, points, expect)
}

func TestSample3(t *testing.T) {
	points := [][]int{
		{33, 38},
		{32, 96},
		{45, 75},
		{10, 35},
		{68, 38},
		{23, 60},
		{10, 20},
		{49, 4},
		{35, 49},
		{97, 61},
	}
	k := 3
	expect := []int{2, 3, 10}
	runSample(t, k, points, expect)
}
