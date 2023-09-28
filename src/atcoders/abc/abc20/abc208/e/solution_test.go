package main

import "testing"

func runSample(t *testing.T, n int64, k int64, expect int64) {
	res := solve(n, k)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 13, 2, 5)
}

func TestSample2(t *testing.T) {
	runSample(t, 100, 80, 99)
}

func TestSample3(t *testing.T) {
	runSample(t, 1000000000000000000, 1000000000, 841103275147365677)
}

func TestSample4(t *testing.T) {
	var n int64 = 1000000
	var k int64 = 999
	expect := bruteForce(n, k)
	runSample(t, n, k, expect)
}
