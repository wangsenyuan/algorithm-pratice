package main

import "testing"

func runSample(t *testing.T, R, C int, G int64) {
	res := solve(R, C, G)

	var sum int64

	for i := 0; i < len(res); i++ {
		sum += res[i]
	}

	if sum != G {
		t.Errorf("Sample %d %d %d, result %v, is wrong", R, C, G, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 3, 2, 5)
}

func TestSample2(t *testing.T) {
	runSample(t, 3, 3, 10)
}

func TestSample3(t *testing.T) {
	runSample(t, 5, 4, 7)
}
