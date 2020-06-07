package main

import "testing"

func runSample(t *testing.T, N int, p []float64, F []int, s float64, expect int) {
	res := solve(N, p, F, s)

	if res != expect {
		t.Errorf("Sample %d %v %v %f, expect %d, but got %d", N, p, F, s, expect, res)
	}
}

func TestSample1(t *testing.T) {
	N := 2
	p := []float64{0.8, 0.9}
	F := []int{100, 200}
	s := 1.0
	expect := 300
	runSample(t, N, p, F, s, expect)
}

func TestSample2(t *testing.T) {
	N := 4
	p := []float64{0.28, 0.62, 0.92, 0.96}
	F := []int{3, 8, 7, 2}
	s := 0.05
	expect := 9
	runSample(t, N, p, F, s, expect)
}
