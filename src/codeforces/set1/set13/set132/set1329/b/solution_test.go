package main

import "testing"

func runSample(t *testing.T, d int, m int, expect int) {
	res := solve(d, m)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 1, 1000000000, 1)
}

func TestSample2(t *testing.T) {
	runSample(t, 2, 999999999, 3)
}

func TestSample3(t *testing.T) {
	runSample(t, 3, 99999998, 5)
}

func TestSample4(t *testing.T) {
	// 1, 2, 4
	// (1, 0), (0,2,3), (4, 0)
	runSample(t, 4, 9999997, 11)
}

func TestSample5(t *testing.T) {
	runSample(t, 5, 999996, 17)
}

func TestSample6(t *testing.T) {
	runSample(t, 6, 99995, 23)
}

func TestSample7(t *testing.T) {
	runSample(t, 7, 9994, 29)
}

func TestSample8(t *testing.T) {
	runSample(t, 8, 993, 59)
}

func TestSample9(t *testing.T) {
	runSample(t, 9, 92, 89)
}

func TestSample10(t *testing.T) {
	runSample(t, 10, 1, 0)
}
