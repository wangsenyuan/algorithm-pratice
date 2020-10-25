package main

import "testing"

func runSample(t *testing.T, n int, h int, expect uint64) {
	res := solve(n, h)

	if res != expect {
		t.Errorf("Sample %d %d, expect %d, but got %d", n, h, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 3, 2, 5)
}

func TestSample2(t *testing.T) {
	runSample(t, 3, 3, 4)
}

func TestSample3(t *testing.T) {
	runSample(t, 27, 11, 61162698256896)
}
//
//func TestSample4(t *testing.T) {
//	runSample(t, 1<<5-1, 5, 1)
//}
