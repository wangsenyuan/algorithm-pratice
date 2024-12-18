package main

import "testing"

func runSample(t *testing.T, n int) {
	res := solve(n)
	var cnt int
	var sum int

	for i := 1; i < len(res); i++ {
		if res[i] == '3' {
			cnt++
		} else {
			// res[i] = 7
			sum += cnt * (cnt - 1) / 2
		}
	}

	if sum != n {
		t.Fatalf("Sample result %v, not correct", res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 10)
}

func TestSample2(t *testing.T) {
	runSample(t, 1)
}

func TestSample3(t *testing.T) {
	runSample(t, 1e9)
}
