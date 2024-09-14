package main

import "testing"

func runSample(t *testing.T, a int, m int, expect int) {
	res := solve(a, m)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	// gcd(4, 9) = 1
	// x = 0,1,3,4,6,7
	// l = 4, r = (4 + 9) / g = 13
	// 4 * 1 - 4, 5 * 1 - 4, 6 * 1 - 4
	// x = 2, gcd(6, 9) = 3

	runSample(t, 4, 9, 6)
}

func TestSample2(t *testing.T) {
	runSample(t, 5, 10, 1)
}
