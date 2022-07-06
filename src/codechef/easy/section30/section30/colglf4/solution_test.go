package main

import "testing"

func runSample(t *testing.T, N int, E int, H int, A, B, C int, expect int64) {
	res := solve(N, E, H, A, B, C)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	N, E, H, A, B, C := 5, 4, 4, 2, 2, 2
	var expect int64 = -1
	runSample(t, N, E, H, A, B, C, expect)
}

func TestSample2(t *testing.T) {
	N, E, H, A, B, C := 4, 5, 5, 1, 2, 3
	var expect int64 = 7
	runSample(t, N, E, H, A, B, C, expect)
}

func TestSample3(t *testing.T) {
	N, E, H, A, B, C := 4, 5, 5, 3, 2, 1
	var expect int64 = 4
	runSample(t, N, E, H, A, B, C, expect)
}
