package main

import "testing"

func runSample(t *testing.T, n int, expect int64) {
	res := solve(n)

	if res != expect {
		t.Errorf("Sample %d expect %d, but got %d", n, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 221, 333)
}

func TestSample2(t *testing.T) {
	runSample(t, 333, 1333)
}

func TestSample3(t *testing.T) {
	runSample(t, 3002, 3033)
}

func TestSample4(t *testing.T) {
	for i := 10; i <= 100000; i++ {
		runSample(t, i, solve1(i))
	}
}

func TestSample5(t *testing.T) {
	runSample(t, 2641, 3033)
}

func TestSample6(t *testing.T) {
	runSample(t, 329, 333)
}

func TestSample7(t *testing.T) {
	runSample(t, 332, 333)
}

func TestSample8(t *testing.T) {
	runSample(t, 29333, 30033)
}

func TestSample9(t *testing.T) {
	runSample(t, 99333, 100333)
}
