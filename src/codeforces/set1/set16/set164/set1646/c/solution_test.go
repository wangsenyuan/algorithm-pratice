package main

import "testing"

func runSample(t *testing.T, n int64, expect int) {
	res := solve(n)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	var n int64 = 7
	expect := 2
	runSample(t, n, expect)
}

func TestSample2(t *testing.T) {
	var n int64 = 11
	expect := 3
	runSample(t, n, expect)
}

func TestSample3(t *testing.T) {
	var n int64 = 240
	expect := 4
	runSample(t, n, expect)
}

func TestSample4(t *testing.T) {
	var n int64 = 17179869184
	expect := 1
	runSample(t, n, expect)
}

func TestSample5(t *testing.T) {
	var n int64 = 378404098055
	expect := 13
	runSample(t, n, expect)
}

func TestSample6(t *testing.T) {
	var n int64 = 432931785912
	expect := 12
	runSample(t, n, expect)
}
