package main

import "testing"

func runSample(t *testing.T, n int, expect int) {
	res := solve(n)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 4, 750000007)
}

func TestSample2(t *testing.T) {
	runSample(t, 2, 500000005)
}

func TestSample3(t *testing.T) {
	runSample(t, 9999999, 412120619)
}

func TestSample4(t *testing.T) {
	runSample(t, 10000000, 111148956)
}
