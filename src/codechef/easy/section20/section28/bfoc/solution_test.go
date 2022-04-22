package main

import "testing"

func runSample(t *testing.T, n int, m int, l int, expect int) {
	res := solve(int64(n), int64(m), int64(l))

	if int(res) != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 2, 8, 2, 0)
}

func TestSample2(t *testing.T) {
	runSample(t, 3, 9, 2, 1)
}

func TestSample3(t *testing.T) {
	runSample(t, 3, 10, 2, 0)
}

func TestSample4(t *testing.T) {
	runSample(t, 4, 100, 10, 9)
}