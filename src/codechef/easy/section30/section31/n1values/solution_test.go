package main

import "testing"

func runSample(t *testing.T, n int) {
	res := solve(n)
	var expect int64 = 1 << uint64(n)
	var sum int64

	for i := 0; i < len(res); i++ {
		sum += res[i]
	}

	if sum != expect {
		t.Errorf("Sample result %v, getting wrong sum %d", res, sum)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 60)
}
