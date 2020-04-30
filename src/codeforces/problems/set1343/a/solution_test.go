package main

import "testing"

func runSample(t *testing.T, n int) {
	res := solve(n)

	if n%res != 0 {
		t.Fatalf("Sample %d, result %d, not correct", n, res)
	}

	y := n / res

	if !isPower2(int64(y + 1)) {
		t.Fatalf("Sample %d, result %d, not correct", n, res)
	}
	t.Logf("sample %d, result is %d", n, res)
}

func TestSample1(t *testing.T) {
	runSample(t, 3)
}

func TestSample2(t *testing.T) {
	runSample(t, 6)
}

func TestSample3(t *testing.T) {
	runSample(t, 7)
}

func TestSample4(t *testing.T) {
	runSample(t, 21)
}

func TestSample5(t *testing.T) {
	runSample(t, 28)
}

func TestSample6(t *testing.T) {
	runSample(t, 999999999)
}

func TestSample7(t *testing.T) {
	runSample(t, 999999984)
}
