package main

import "testing"

func runSample(t *testing.T, N, L, A, B int, M string, expect int) {
	res := solve(N, L, A, B, M)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	N, L, A, B := 6, 6, 0, 0
	M := "+-+-+-"
	expect := 5
	runSample(t, N, L, A, B, M, expect)
}

func TestSample2(t *testing.T) {
	N, L, A, B := 6, 6, 3, 3
	M := "+-+-+-"
	expect := 9
	runSample(t, N, L, A, B, M, expect)
}

func TestSample3(t *testing.T) {
	N, L, A, B := 6, 1, 0, 0
	M := "+-+-+-"
	expect := 4
	runSample(t, N, L, A, B, M, expect)
}
func TestSample4(t *testing.T) {
	N, L, A, B := 6, 6, 1, 4
	M := "+-+-+-"
	expect := 1
	runSample(t, N, L, A, B, M, expect)
}

func TestSample5(t *testing.T) {
	N, L, A, B := 6, 6, 2, 6
	M := "+-+-+-"
	expect := 0
	runSample(t, N, L, A, B, M, expect)
}