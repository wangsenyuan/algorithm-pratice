package main

import "testing"

func runSample(t *testing.T, x0, y0 int, ax int, ay int, bx int, by int, xs int, ys int, w int, expect int) {
	res := solve(x0, y0, ax, ay, bx, by, xs, ys, w)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	x0, y0, ax, ay, bx, by := 1, 1, 2, 3, 1, 0
	xs, ys, w := 2, 4, 20
	expect := 3
	runSample(t, x0, y0, ax, ay, bx, by, xs, ys, w, expect)
}

func TestSample2(t *testing.T) {
	x0, y0, ax, ay, bx, by := 1, 1, 2, 3, 1, 0
	xs, ys, w := 15, 27, 26
	expect := 2
	runSample(t, x0, y0, ax, ay, bx, by, xs, ys, w, expect)
}

func TestSample3(t *testing.T) {
	x0, y0, ax, ay, bx, by := 1, 1, 2, 3, 1, 0
	xs, ys, w := 2, 2, 1
	expect := 0
	runSample(t, x0, y0, ax, ay, bx, by, xs, ys, w, expect)
}

func TestSample4(t *testing.T) {
	x0, y0, ax, ay, bx, by := 1, 1, 10, 10, 1, 2
	xs, ys, w := 10000000000000000, 10000000000000000, 10000000000000000
	expect := 1
	runSample(t, x0, y0, ax, ay, bx, by, xs, ys, w, expect)
}

func TestSample5(t *testing.T) {
	x0, y0, ax, ay, bx, by := 10000000000000000, 10000000000000000, 100, 100, 10000000000000000, 10000000000000000
	xs, ys, w := 10000000000000000, 10000000000000000, 10000000000000000
	expect := 1
	runSample(t, x0, y0, ax, ay, bx, by, xs, ys, w, expect)
}

func TestSample6(t *testing.T) {
	x0, y0, ax, ay, bx, by := 923, 247, 2, 2, 1, 1
	xs, ys, w := 1000000000, 1000000000, 353416061
	expect := 0
	runSample(t, x0, y0, ax, ay, bx, by, xs, ys, w, expect)
}

func TestSample7(t *testing.T) {
	x0, y0, ax, ay, bx, by := 562949953421311, 562949953421311, 32, 32, 999, 999
	xs, ys, w := 1023175, 1023175, 1
	expect := 0
	runSample(t, x0, y0, ax, ay, bx, by, xs, ys, w, expect)
}
