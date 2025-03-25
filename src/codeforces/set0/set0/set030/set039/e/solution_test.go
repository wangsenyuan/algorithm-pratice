package main

import "testing"

func runSample(t *testing.T, a, b, n int, expect string) {
	res := solve(a, b, n)
	if res != expect {
		t.Errorf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 2, 2, 10, "Masha")
}

func TestSample2(t *testing.T) {
	runSample(t, 5, 5, 16808, "Masha")
}

func TestSample3(t *testing.T) {
	runSample(t, 3, 1, 4, "Stas")
}

func TestSample4(t *testing.T) {
	runSample(t, 1, 4, 10, "Missing")
}

func TestSample5(t *testing.T) {
	runSample(t, 1, 3, 60, "Masha")
}

func TestSample6(t *testing.T) {
	runSample(t, 3, 1, 64, "Masha")
}
