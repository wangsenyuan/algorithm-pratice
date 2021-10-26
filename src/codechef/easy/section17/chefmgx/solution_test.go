package main

import "testing"

func runSample(t *testing.T, X, Y int, expect int) {
	res := solve(X, Y)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 2, 10, 6)
}

func TestSample2(t *testing.T) {
	runSample(t, 5, 12, 5)
}

func TestSample3(t *testing.T) {
	runSample(t, 34, 43, 6)
}

func TestSample4(t *testing.T) {
	runSample(t, 57, 63, 4)
}
func TestSample5(t *testing.T) {
	runSample(t, 2, 3, 1)
}
