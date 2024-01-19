package main

import "testing"

func runSample(t *testing.T, n int, expect int) {
	res := solve(n)

	if res != expect {
		t.Fatalf("Sample expect %d, but guto %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 2, 2)
}

func TestSample2(t *testing.T) {
	runSample(t, 4, 17)
}

func TestSample3(t *testing.T) {
	runSample(t, 3, 7)
}

func TestSample4(t *testing.T) {
	runSample(t, 10, 303)
}

func TestSample5(t *testing.T) {
	runSample(t, 20, 2529)
}
