package main

import "testing"

func runSample(t *testing.T, k int, expect int) {
	res := solve(k)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 1, 0)
}


func TestSample2(t *testing.T) {
	runSample(t, 11, 1)
}

func TestSample3(t *testing.T) {
	runSample(t, 123456, 127)
}

func TestSample4(t *testing.T) {
	runSample(t, 2, 0)
}