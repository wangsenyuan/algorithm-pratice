package main

import "testing"

func runSample(t *testing.T, n int, m int, x1 int, y1 int, x2 int, y2 int, expect int) {
	res := solve(n, m, x1, y1, x2, y2)

	if int(res) != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, m, x1, y1, x2, y2 := 4, 4, 1, 1, 3, 3
	expect := 8
	runSample(t, n, m, x1, y1, x2, y2, expect)
}
func TestSample2(t *testing.T) {
	n, m, x1, y1, x2, y2 := 4, 3, 1, 1, 2, 2
	expect := 2
	runSample(t, n, m, x1, y1, x2, y2, expect)
}

func TestSample3(t *testing.T) {
	n, m, x1, y1, x2, y2 := 3, 3, 3, 2, 1, 1
	expect := 5
	runSample(t, n, m, x1, y1, x2, y2, expect)
}

func TestSample4(t *testing.T) {
	n, m, x1, y1, x2, y2 := 2, 5, 1, 5, 2, 2
	expect := 6
	runSample(t, n, m, x1, y1, x2, y2, expect)
}

// 100 100 99 1 1 99

func TestSample5(t *testing.T) {
	n, m, x1, y1, x2, y2 := 100, 100, 99, 1, 1, 99
	expect := 9992
	runSample(t, n, m, x1, y1, x2, y2, expect)
}
