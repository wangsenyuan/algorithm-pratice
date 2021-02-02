package main

import "testing"

func runSample(t *testing.T, S string, expect int) {
	res := solve(len(S), []byte(S))

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "01", 0)
}

func TestSample2(t *testing.T) {
	runSample(t, "110", 1)
}

func TestSample3(t *testing.T) {
	runSample(t, "10101", 2)
}

func TestSample4(t *testing.T) {
	runSample(t, "10011", 1)
}

func TestSample5(t *testing.T) {
	runSample(t, "111000", 3)
}

func TestSample6(t *testing.T) {
	runSample(t, "1101000010110111010", 7)
}
