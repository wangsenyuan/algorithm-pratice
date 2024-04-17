package main

import "testing"

func runSample(t *testing.T, L int, R int, expect int) {
	res := solve(L, R)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 1, 1000, 1000)
}

func TestSample2(t *testing.T) {
	runSample(t, 1024, 1024, 1)
}

func TestSample3(t *testing.T) {
	runSample(t, 65536, 65536, 0)
}

func TestSample4(t *testing.T) {
	runSample(t, 999999, 1000001, 2)
}
