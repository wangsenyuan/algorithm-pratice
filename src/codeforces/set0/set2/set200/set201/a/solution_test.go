package main

import "testing"

func runSample(t *testing.T, x int, expect int) {
	res := solve(x)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 1, 1)
}

func TestSample2(t *testing.T) {
	runSample(t, 4, 3)
}


func TestSample3(t *testing.T) {
	runSample(t, 9, 5)
}

func TestSample4(t *testing.T) {
	runSample(t, 3, 5)
}