package main

import "testing"

func runSample(t *testing.T, r int, expect int) {
	res := solve(r)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 1, 8)
}

func TestSample2(t *testing.T) {
	runSample(t, 2, 16)
}
func TestSample3(t *testing.T) {
	runSample(t, 3, 20)
}

func TestSample4(t *testing.T) {
	runSample(t, 4, 24)
}

func TestSample5(t *testing.T) {
	runSample(t, 5, 40)
}

func TestSample6(t *testing.T) {
	runSample(t, 1984, 12504)
}
