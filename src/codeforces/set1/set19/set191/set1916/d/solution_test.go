package main

import "testing"

func runSample(t *testing.T, n int) {
	res := solve(n)

	if len(res) != n {
		t.Fatalf("Sample result %v, not correct", res)
	}

	for _, ans := range res {
		if len(ans) != n {
			t.Fatalf("Sample result %v, not correct", res)
		}
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 13)
}
