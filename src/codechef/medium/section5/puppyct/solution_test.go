package main

import "testing"

func runSample(t *testing.T, N, K int, dumps []int, expect int64) {
	res := solve(N, K, dumps)
	if res != expect {
		t.Errorf("Sample %d %d %v, expect %d, but got %d", N, K, dumps, expect, res)
	}
}

func TestSample1(t *testing.T) {
	N, K := 3, 2
	dumps := []int{
		1, 2,
		3, 2,
	}
	expect := 5
	runSample(t, N, K, dumps, int64(expect))
}

func TestSample2(t *testing.T) {
	N, K := 4, 2
	dumps := []int{
		2, 2,
		3, 3,
	}
	expect := 8
	runSample(t, N, K, dumps, int64(expect))
}
