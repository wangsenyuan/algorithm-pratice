package main

import "testing"

func runSample(t *testing.T, N, M, L, R int, expect uint64) {
	res := solve(N, M, L, R)
	if res != expect {
		t.Errorf("Sample %d %d %d %d, expect %d, but got %d", N, M, L, R, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 5, 5, 1, 5, 3125)
}

func TestSample2(t *testing.T) {
	runSample(t, 5, 5, 4, 5, 2)
}

func runFastPow(t *testing.T, a, n int, expect uint64) {
	res := fastPow(a, n)

	if res != expect {
		t.Errorf("fastPow(%d, %d), expect %d, but got %d", a, n, expect, res)
	}
}

func TestFastPow1(t *testing.T) {
	runFastPow(t, 2, 3, 8)
}

func TestFastPow2(t *testing.T) {
	runFastPow(t, 3, 2, 9)
}

func TestFastPow3(t *testing.T) {
	runFastPow(t, 3, 1, 3)
}

func TestFastPow4(t *testing.T) {
	runFastPow(t, 5, 4, 625)
}
