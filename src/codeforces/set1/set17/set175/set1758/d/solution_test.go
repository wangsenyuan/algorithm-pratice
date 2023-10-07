package main

import "testing"

func runSample(t *testing.T, n int) {
	res := solve(n)

	var sum int
	for i := 0; i < n; i++ {
		sum += res[i]
	}
	diff := res[n-1] - res[0]
	if diff*diff != sum {
		t.Fatalf("Sample result %v, not correct", res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 2)
}

func TestSample2(t *testing.T) {
	runSample(t, 5)
}

func TestSample3(t *testing.T) {
	runSample(t, 4)
}
