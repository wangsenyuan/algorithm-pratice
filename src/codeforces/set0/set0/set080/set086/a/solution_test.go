package main

import "testing"

func runSample(t *testing.T, l int, r int, expect int) {
	res := solve(l, r)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 3, 7, 20)
}

func TestSample2(t *testing.T) {
	runSample(t, 1, 1, 8)
}

func TestSample3(t *testing.T) {
	runSample(t, 8, 10, 890)
}

func TestSample5(t *testing.T) {
	runSample(t, 1, 999, 249500)
}
