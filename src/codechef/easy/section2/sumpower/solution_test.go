package main

import "testing"

func runSample(t *testing.T, N, K int, S string, expect int) {
	res := solve(N, K, S)
	if res != expect {
		t.Errorf("Sample %d %d %s, expect %d, but got %d", N, K, S, expect, res)
	}
}

func TestSample1(t *testing.T) {
	N, K := 6, 3
	S := "aabbcc"
	expect := 4
	runSample(t, N, K, S, expect)
}

func TestSample2(t *testing.T) {
	N, K := 5, 2
	S := "abccc"
	expect := 3
	runSample(t, N, K, S, expect)
}

func TestSample3(t *testing.T) {
	N, K := 4, 3
	S := "aabb"
	expect := 1
	runSample(t, N, K, S, expect)
}
