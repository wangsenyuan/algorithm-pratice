package main

import (
	"sort"
	"testing"
)

func runSample(t *testing.T, pts [][]int) {
	res := solve(pts)

	n := len(pts)

	points := make([]Point, n)

	for i := 0; i < n; i++ {
		points[i] = Point{i, pts[i][0], pts[i][1]}
	}

	sort.Slice(points, func(i, j int) bool {
		return points[i].x < points[j].x || points[i].x == points[j].x && points[i].y < points[j].y
	})

	for i := 0; i < n; {
		j := i
		var diff int
		for i < n && points[i].x == points[j].x {
			if res[points[i].id] == 'r' {
				diff++
			} else {
				diff--
			}
			i++
		}
		if abs(diff) > 1 {
			t.Fatalf("Sample result %v, not correct", res)
		}
	}

	sort.Slice(points, func(i, j int) bool {
		return points[i].y < points[j].y || points[i].y == points[j].y && points[i].x < points[j].x
	})

	for i := 0; i < n; {
		j := i
		var diff int
		for i < n && points[i].y == points[j].y {
			if res[points[i].id] == 'r' {
				diff++
			} else {
				diff--
			}
			i++
		}
		if abs(diff) > 1 {
			t.Fatalf("Sample result %v, not correct", res)
		}
	}
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func TestSample1(t *testing.T) {
	pts := [][]int{
		{1, 1},
		{1, 2},
		{2, 1},
		{2, 2},
	}
	runSample(t, pts)
}

func TestSample2(t *testing.T) {
	pts := [][]int{
		{1, 1},
		{1, 2},
		{2, 1},
	}
	runSample(t, pts)
}
