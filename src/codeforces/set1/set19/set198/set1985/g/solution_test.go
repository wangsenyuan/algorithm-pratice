package main

import "testing"

func runSample(t *testing.T, l int, r int, k int, expect int) {
	res := solve(l, r, k)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 0, 1, 4, 2)
}

func TestSample2(t *testing.T) {
	runSample(t, 0, 2, 7, 3)
}
func TestSample3(t *testing.T) {
	runSample(t, 1, 2, 1, 90)
}
func TestSample4(t *testing.T) {
	runSample(t, 1, 2, 3, 12)
}

func TestSample5(t *testing.T) {
	runSample(t, 582, 74663, 3, 974995667)
}

func TestSample6(t *testing.T) {
	runSample(t, 0, 3, 1, 999)
}