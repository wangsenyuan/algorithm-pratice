package main

import "testing"

func runSample(t *testing.T, x int, y int, expect int) {
	res := solve(x, y)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	x, y := 3, 9
	expect := 3
	runSample(t, x, y, expect)
}

func TestSample2(t *testing.T) {
	x, y := 5, 8
	expect := 0
	runSample(t, x, y, expect)
}

func TestSample3(t *testing.T) {
	x, y := 2, 12
	expect := 27
	runSample(t, x, y, expect)
}

func TestSample4(t *testing.T) {
	x, y := 1, 1000000000

	expect := 824916815
	runSample(t, x, y, expect)
}

func TestSample5(t *testing.T) {
	x, y := 495219, 444706662

	expect := 115165527
	runSample(t, x, y, expect)
}
