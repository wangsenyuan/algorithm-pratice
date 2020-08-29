package main

import "testing"

func runSample(t *testing.T, n int, s string, expect string) {
	res := solve(n, []byte(s))

	if res != expect {
		t.Errorf("Sample %d %s, expect %s, but got %s", n, s, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 10, "0001111111", "0001111111")
}

func TestSample2(t *testing.T) {
	runSample(t, 4, "0101", "001")
}

func TestSample3(t *testing.T) {
	runSample(t, 8, "11001101", "01")
}

func TestSample4(t *testing.T) {
	runSample(t, 10, "1110000000", "0")
}

func TestSample5(t *testing.T) {
	runSample(t, 1, "1", "1")
}
