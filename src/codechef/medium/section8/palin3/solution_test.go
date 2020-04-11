package main

import "testing"

func runSample(t *testing.T, s string, expect int64) {
	res := solve(s)

	if res != expect {
		t.Errorf("Sample %s, expect %d, but got %d", s, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "1045003", 4)
}

func TestSample2(t *testing.T) {
	runSample(t, "303", 4)
}

func TestSample3(t *testing.T) {
	runSample(t, "3003", 5)
}

func TestSample4(t *testing.T) {
	runSample(t, "21300312", 6)
}

func TestSample5(t *testing.T) {
	runSample(t, "2300312", 5)
}

func TestSample6(t *testing.T) {
	runSample(t, "414", 1)
}
