package main

import "testing"

func runSample(t *testing.T, l int, r int, expect int) {
	res := solve(l, r)
	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 1, 3, 7)
}

func TestSample2(t *testing.T) {
	// 1, 3,5,7,9
	// 2, 4,6,8,10,12,14,16,18
	runSample(t, 5, 14, 105)
}

func TestSample3(t *testing.T) {
	// 1, 3,5,7,9
	// 2, 4,6,8,10,12,14,16,18
	runSample(t, 88005553535, 99999999999, 761141116)
}

func TestSample4(t *testing.T) {
	runSample(t, 2, 3, 6)
}
