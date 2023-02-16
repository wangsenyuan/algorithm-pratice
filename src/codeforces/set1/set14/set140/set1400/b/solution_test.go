package main

import "testing"

func runSample(t *testing.T, p, f int, cnts int, cntw int, s, w int, expect int64) {
	res := solve(p, f, cnts, cntw, s, w)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	p, f := 33, 27
	cnts, cntw := 6, 10
	s, w := 5, 6
	var expect int64 = 11
	runSample(t, p, f, cnts, cntw, s, w, expect)
}

func TestSample2(t *testing.T) {
	p, f := 100, 200
	cnts, cntw := 10, 10
	s, w := 5, 5
	var expect int64 = 20
	runSample(t, p, f, cnts, cntw, s, w, expect)
}

func TestSample3(t *testing.T) {
	p, f := 1, 19
	cnts, cntw := 1, 3
	s, w := 19, 5
	var expect int64 = 3
	runSample(t, p, f, cnts, cntw, s, w, expect)
}

func TestSample4(t *testing.T) {
	p, f := 28, 94
	cnts, cntw := 14, 8
	s, w := 6, 5
	var expect int64 = 21
	runSample(t, p, f, cnts, cntw, s, w, expect)
}
