package main

import "testing"

func runSample(t *testing.T, n int, expect int) {
	res := solve(n)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 6, 1)
}

func TestSample2(t *testing.T) {
	runSample(t, 9, 1)
}

func TestSample3(t *testing.T) {
	runSample(t, 8070, 63)
}
func TestSample4(t *testing.T) {
	runSample(t, 3612, 41)
}

func TestSample5(t *testing.T) {
	runSample(t, 5, 1)
}
