package main

import "testing"

func runSample(t *testing.T, x int, n int, expect int) {
	res := solve(x, n)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 10, 2, 2)
}

func TestSample2(t *testing.T) {
	runSample(t, 20190929, 1605, 363165664)
}

func TestSample3(t *testing.T) {
	// g(1, 2) = 1
	// g(2, 2) = 2
	// g(3, 2) = 1
	// g(4, 2) = 4
	// g(5, 2) = 1
	// g(6, 2) = 2
	// g(7, 2) = 1
	// g(8, 2) = 8
	// g(5, 5) = 5
	runSample(t, 10, 8, 16*8*5)
}

func TestSample4(t *testing.T) {
	runSample(t, 947, 987654321987654321, 593574252)
}
