package main

import "testing"

func runSample(t *testing.T, x1, y1, r1, x2, y2, r2 int, expect []float64) {
	res := solve(x1, y1, r1, x2, y2, r2)

	if len(expect) == 0 {
		if len(res) != 0 {
			t.Errorf("Sample %d %d %d %d %d %d, expect %v, but got %v", x1, y1, r1, x2, y2, r2, expect, res)
		}
	} else {
		if len(res) == 0 {
			t.Errorf("Sample %d %d %d %d %d %d, expect %v, but got %v", x1, y1, r1, x2, y2, r2, expect, res)
		} else if !inCycle(x1, y1, r1, res) || inCycle(x2, y2, r2, res) {
			t.Errorf("Sample %d %d %d %d %d %d, expect %v, but got %v", x1, y1, r1, x2, y2, r2, expect, res)
		}
	}
}

func inCycle(x, y, r int, point []float64) bool {
	X := float64(x)
	Y := float64(y)
	R := float64(r)

	dx := (X - point[0])
	dy := (Y - point[1])
	dd := dx*dx + dy*dy

	return dd <= R*R
}

func TestSample1(t *testing.T) {
	x1, y1, r1 := 0, 0, 2
	x2, y2, r2 := 0, 0, 3
	runSample(t, x1, y1, r1, x2, y2, r2, nil)
}

func TestSample2(t *testing.T) {
	x1, y1, r1 := 5, 5, 1
	x2, y2, r2 := 0, 0, 1
	expect := []float64{5.0, 5.0}
	runSample(t, x1, y1, r1, x2, y2, r2, expect)
}
func TestSample3(t *testing.T) {
	x1, y1, r1 := 0, 1, 1
	x2, y2, r2 := 0, 2, 2
	runSample(t, x1, y1, r1, x2, y2, r2, nil)
}
