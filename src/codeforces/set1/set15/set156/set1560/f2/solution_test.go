package main

import "testing"

func runSample(t *testing.T, n int, k int, expect int) {
	res := solve(n, k)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 1, 1, 1)
}

func TestSample2(t *testing.T) {
	runSample(t, 221, 2, 221)
}
func TestSample3(t *testing.T) {
	runSample(t, 177890, 2, 181111)
}

func TestSample4(t *testing.T) {
	runSample(t, 998244353, 1, 999999999)
}

func TestSample5(t *testing.T) {
	runSample(t, 109, 2, 110)
}

func TestSample6(t *testing.T) {
	runSample(t, 19, 1, 22)
}
