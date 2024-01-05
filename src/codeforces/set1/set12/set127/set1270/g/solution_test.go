package main

import "testing"

func runSample(t *testing.T, a []int32) {
	res := solve(a)

	var sum int32

	for _, i := range res {
		sum += a[i-1]
	}

	if len(res) == 0 || sum != 0 {
		t.Fatalf("Sample result %v, not correct", res)
	}
}

func TestSample1(t *testing.T) {
	a := []int32{0, 1, 2, 3, 4}
	runSample(t, a)
}

func TestSample2(t *testing.T) {
	a := []int32{-3, 1, 1, 1}
	runSample(t, a)
}
