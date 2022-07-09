package main

import "testing"

func runSample(t *testing.T, A []int, expect int64) {
	res := solve(A)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{500, 500}
	var expect int64 = 250000

	runSample(t, A, expect)
}

func TestSample2(t *testing.T) {
	A := []int{500, 500, 500}
	var expect int64 = 500000

	runSample(t, A, expect)
}

func TestSample3(t *testing.T) {
	A := []int{100, 800, 300, 500}
	var expect int64 = 2080000

	runSample(t, A, expect)
}

func TestSample4(t *testing.T) {
	A := []int{300, 700, 800, 200}
	var expect int64 = 2250000

	runSample(t, A, expect)
}
