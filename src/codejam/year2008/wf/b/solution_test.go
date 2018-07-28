package main

import "testing"

func runSample(t *testing.T, W, H int, dx1 int, dy1 int, dx2 int, dy2 int, x, y int, expect int64) {
	res := solve(W, H, dx1, dy1, dx2, dy2, x, y)

	if res != expect {
		t.Errorf("sample %d %d %d %d %d %d %d %d, expect %d, but got %d", W, H, dx1, dy1, dx2, dy2, x, y, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 5, 3, -1, 0, -1, -1, 4, 2, 12)
}

func TestSample2(t *testing.T) {
	runSample(t, 50, 50, 0, 1, 1, 1, 10, 10, 820)
}

func TestSample3(t *testing.T) {
	runSample(t, 6, 2, 2, 0, 3, 0, 0, 0, 5)
}

func TestSample4(t *testing.T) {
	runSample(t, 86, 92, 1, 0, -1, 0, 42, 46, 86)
}
