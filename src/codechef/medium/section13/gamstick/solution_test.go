package main

import "testing"

func runSample(t *testing.T, n, x1, y1, x2, y2 uint64, expect string) {
	res := solve(n, x1, y1, x2, y2)

	if res != expect {
		t.Errorf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	var n, x1, y1, x2, y2 uint64 = 5, 1, 1, 2, 3
	expect := Slava
	runSample(t, n, x1, y1, x2, y2, expect)
}

func TestSample2(t *testing.T) {
	var n, x1, y1, x2, y2 uint64 = 1, 1, 1, 2, 1
	expect := Draw
	runSample(t, n, x1, y1, x2, y2, expect)
}

func TestSample3(t *testing.T) {
	var n, x1, y1, x2, y2 uint64 = 3, 1, 1, 2, 3
	expect := Draw
	runSample(t, n, x1, y1, x2, y2, expect)
}

func TestSample4(t *testing.T) {
	var n, x1, y1, x2, y2 uint64 = 4, 1, 2, 1, 4
	expect := Miron
	runSample(t, n, x1, y1, x2, y2, expect)
}

func TestSample5(t *testing.T) {
	var n, x1, y1, x2, y2 uint64 = 4, 1, 2, 2, 4
	expect := Draw
	runSample(t, n, x1, y1, x2, y2, expect)
}
