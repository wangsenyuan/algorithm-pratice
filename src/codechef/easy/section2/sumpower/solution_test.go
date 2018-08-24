package main

import "testing"

func runSample(t *testing.T, N, K int, S string, expect int64) {
	res := solve(N, K, S)
	if res != expect {
		t.Errorf("Sample %d %d %s, expect %d, but got %d", N, K, S, expect, res)
	}
}

func TestSample1(t *testing.T) {
	N, K := 6, 3
	S := "aabbcc"
	runSample(t, N, K, S, 4)
}

func TestSample2(t *testing.T) {
	N, K := 5, 2
	S := "abccc"
	runSample(t, N, K, S, 3)
}

func TestSample3(t *testing.T) {
	N, K := 4, 3
	S := "aabb"
	runSample(t, N, K, S, 1)
}
