package main

import "testing"

func runSample(t *testing.T, n int, m int, expect int) {
	res := solve(n, m)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 1, 2, 2)
}

func TestSample2(t *testing.T) {
	runSample(t, 2, 3, 5)
}

func TestSample3(t *testing.T) {
	runSample(t, 3, 6, 10)
}

func TestSample4(t *testing.T) {
	runSample(t, 10, 1024, 2046)
}

func TestSample5(t *testing.T) {
	runSample(t, 11, 550, 408)
}
