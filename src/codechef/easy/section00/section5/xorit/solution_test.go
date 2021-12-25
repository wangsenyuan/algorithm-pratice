package main

import "testing"

func runSample(t *testing.T, L int, R int, expect int64) {
	res := solve(L, R)

	if res != expect {
		t.Errorf("Sample %d %d, expect %d, but got %d", L, R, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 1, 10, 28)
}

func TestSample2(t *testing.T) {
	runSample(t, 8, 8, -1)
}

func TestSample3(t *testing.T) {
	runSample(t, 9, 9, 8)
}

func TestSample4(t *testing.T) {
	runSample(t, 3, 6, 9)
}

func TestSample5(t *testing.T) {
	runSample(t, 3, 6, 9)
}

func TestSample6(t *testing.T) {
	runSample(t, 4, 10, 28)
}

func TestSample7(t *testing.T) {
	runSample(t, 10, 17, 79)
}

func TestSample8(t *testing.T) {
	runSample(t, 100, 159, 7485)
}
