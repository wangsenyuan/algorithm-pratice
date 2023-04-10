package main

import "testing"

func runSample(t *testing.T, x int, d int, expect bool) {
	res := solve(x, d)
	if res != expect {
		t.Fatalf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	x, d := 6, 2
	expect := false
	runSample(t, x, d, expect)
}

func TestSample2(t *testing.T) {
	x, d := 12, 2
	expect := false
	runSample(t, x, d, expect)
}

func TestSample3(t *testing.T) {
	x, d := 36, 2
	expect := true
	runSample(t, x, d, expect)
}

func TestSample4(t *testing.T) {
	x, d := 8, 2
	expect := false
	runSample(t, x, d, expect)
}

func TestSample5(t *testing.T) {
	x, d := 1000, 10
	expect := true
	runSample(t, x, d, expect)
}

func TestSample6(t *testing.T) {
	x, d := 2376, 6
	expect := true
	runSample(t, x, d, expect)
}

func TestSample7(t *testing.T) {
	x, d := 128, 4
	expect := false
	runSample(t, x, d, expect)
}

func TestSample8(t *testing.T) {
	x, d := 16384, 4
	expect := true
	runSample(t, x, d, expect)
}

func TestSample9(t *testing.T) {
	x, d := 410338673, 289
	expect := false
	runSample(t, x, d, expect)
}
