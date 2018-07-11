package main

import "testing"

func runSample(t *testing.T, N, K int, A, B string, expect int) {
	res := solve(N, K, A, B)

	if res != expect {
		t.Errorf("Sample %d %d %s %s, expect %d, but got %d", N, K, A, B, expect, res)
	}
}

func TestSample1(t *testing.T) {
	N, K := 3, 2
	A := "aba"
	B := "abb"
	expect := 14
	runSample(t, N, K, A, B, expect)
}

func TestSample2(t *testing.T) {
	N, K := 3, 2
	A := "abc"
	B := "def"
	expect := 13
	runSample(t, N, K, A, B, expect)
}

func TestSample3(t *testing.T) {
	N, K := 1, 1
	A := "a"
	B := "a"
	expect := 1
	runSample(t, N, K, A, B, expect)
}
