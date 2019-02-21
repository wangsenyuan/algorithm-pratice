package main

import "testing"

func runSample(t *testing.T, x1, y1, x2, y2 int) {
	query := func(x, y int) int {
		var dx int
		if x <= x1 {
			dx = x1 - x
		} else if x2 <= x {
			dx = x - x2
		}

		var dy int

		if y <= y1 {
			dy = y1 - y
		} else if y2 <= y {
			dy = y - y2
		}
		return dx + dy
	}

	a, b, c, d := solve(query)
	if a != x1 || b != y1 || c != x2 || d != y2 {
		t.Errorf("Sample %d %d %d %d, result %d %d %d %d not match", x1, y1, x2, y2, a, b, c, d)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 1, 2, 3, 4)
}
