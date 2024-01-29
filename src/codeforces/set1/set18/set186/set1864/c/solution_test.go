package main

import "testing"

func runSample(t *testing.T, x int) {
	res := solve(x)
	if len(res) > 1000 {
		t.Fatalf("Sample took too much for %v", res)
	}

	for i := 0; i+1 < len(res); i++ {
		diff := res[i] - res[i+1]
		if res[i]%diff != 0 {
			t.Fatalf("Sample result %v, not correct", res)
		}
	}

	if res[0] != x || res[len(res)-1] != 1 {
		t.Fatalf("Sample result %v, not correct", res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 14)
}

func TestSample2(t *testing.T) {
	runSample(t, 1e9)
}
