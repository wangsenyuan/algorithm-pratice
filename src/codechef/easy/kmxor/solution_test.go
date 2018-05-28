package main

import "testing"

func runSample(t *testing.T, N, K int, MAX int) {
	res := solve(N, K)

	var m int
	for k := K; k > 0; k >>= 1 {
		m++
	}

	sum := res[0]
	for i := 1; i < N; i++ {
		sum ^= res[i]
		if res[i] == 0 || res[i] > K {
			t.Errorf("Sample %d %d, result %v, have wrong element at %d", N, K, res, i)
			return
		}
	}
	if sum != MAX {
		t.Errorf("Sample %d %d, result %v, xor sum %d is not the max %d", N, K, res, sum, MAX)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 3, 5, 7)
}

func TestSample2(t *testing.T) {
	runSample(t, 4, 3, 3)
}

func TestSample3(t *testing.T) {
	runSample(t, 1, 5, 5)
}

func TestSample4(t *testing.T) {
	runSample(t, 2, 5, 7)
}

func TestSample5(t *testing.T) {
	runSample(t, 4, 1, 0)
}

func TestSample6(t *testing.T) {
	runSample(t, 5, 1, 1)
}
