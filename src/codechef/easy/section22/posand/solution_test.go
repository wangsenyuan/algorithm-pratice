package main

import "testing"

func runSample(t *testing.T, n int, expect bool) {
	res := solve(n)

	if expect != (len(res) == n) {
		t.Errorf("Sample %d expect %t, but got %v", n, expect, res)
		return
	}

	if expect {
		for i := 0; i < n-1; i++ {
			if res[i]&res[i+1] == 0 {
				t.Errorf("Sample %d result %v not correct at %d", n, res, i)
			}
		}
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 4, false)
}

func TestSample2(t *testing.T) {
	for i := 1; i < 100; i++ {
		runSample(t, i, i&(i-1) > 0)
	}
}
