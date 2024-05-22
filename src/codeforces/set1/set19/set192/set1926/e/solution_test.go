package main

import "testing"

func runSample(t *testing.T, n int, k int, expect int) {
	res := solve(n, k)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 7, 1, 1)
}

func TestSample2(t *testing.T) {
	runSample(t, 7, 2, 3)
}

func TestSample3(t *testing.T) {
	runSample(t, 7, 3, 5)
}

func TestSample4(t *testing.T) {
	runSample(t, 7, 4, 7)
}
func TestSample5(t *testing.T) {
	runSample(t, 7, 5, 2)
}

func TestSample6(t *testing.T) {
	runSample(t, 7, 6, 6)
}

func TestSample7(t *testing.T) {
	// 1, 3, 5, 7
	// 2  6
	// 3
	// 4
	runSample(t, 7, 7, 4)
}

func TestSample8(t *testing.T) {
	runSample(t, 34, 14, 27)
}

func TestSample9(t *testing.T) {
	runSample(t, 84, 19, 37)
}

func TestSample10(t *testing.T) {
	runSample(t, 1000000000, 1000000000, 536870912)
}
