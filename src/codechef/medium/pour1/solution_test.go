package main

import "testing"

func runSample(t *testing.T, A, B, C int, expect int) {
	res := solve(A, B, C)
	if res != expect {
		t.Errorf("sample %d %d %d, expect %d, but got %d", A, B, C, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 5, 2, 3, 2)
}

func TestSample2(t *testing.T) {
	runSample(t, 2, 3, 4, -1)
}

func TestSample3(t *testing.T) {
	runSample(t, 2, 3, 1, 2)
}

func TestSample4(t *testing.T) {
	runSample(t, 2, 4, 3, -1)
}
