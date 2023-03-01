package main

import "testing"

func runSample(t *testing.T, a, b, x int64, expect bool) {
	res := solve(a, b, x)

	if res != expect {
		t.Errorf("Sample expect %t, but got %t", expect, res)
	}
}
func TestSample1(t *testing.T) {
	runSample(t, 6, 9, 3, true)
}

func TestSample2(t *testing.T) {
	runSample(t, 15, 38, 7, true)
}

func TestSample3(t *testing.T) {
	runSample(t, 18, 8, 8, true)
}

func TestSample4(t *testing.T) {
	runSample(t, 30, 30, 30, true)
}

func TestSample5(t *testing.T) {
	runSample(t, 40, 50, 90, false)
}

func TestSample6(t *testing.T) {
	runSample(t, 24, 28, 20, true)
}

func TestSample7(t *testing.T) {
	runSample(t, 365, 216, 52, true)
}

func TestSample8(t *testing.T) {
	runSample(t, 537037812705867558, 338887693834423551, 3199921013340, true)
}
