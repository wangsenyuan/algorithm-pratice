package main

import "testing"

func runSample(t *testing.T, n, m int, A, B string, expect int) {
	res := solve(n, m, A, B)

	if res != expect {
		t.Errorf("Sample %d %d %v %v, expect %d, but got %d", n, m, A, B, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 2, 1, "A", "B", 576)
}

func TestSample2(t *testing.T) {
	runSample(t, 4, 3, "ABA", "BBA", 456723)
}
