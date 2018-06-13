package main

import "testing"

func runSample(t *testing.T, A, B, C int64, expect int64) {
	res := solve(A, B, C)
	if res != expect {
		t.Errorf("sample %d %d %d, expect %d, but got %d", A, B, C, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, -5, 0, 5, 0)
}

func TestSample2(t *testing.T) {
	runSample(t, -5, 7, 6, 7)
}

func TestSample3(t *testing.T) {
	runSample(t, -10, -100, 20, 105)
}

func TestSample4(t *testing.T) {
	runSample(t, 1, -1, 1, 2)
}

func TestSample5(t *testing.T) {
	runSample(t, 51, 23, 10, 8)
}
