package main

import "testing"

func runSample(t *testing.T, n, m int, expect int) {
	res := solve(n, m)
	if res != expect {
		t.Errorf("sample %d %d, expect %d, but got %d", n, m, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 3, 8, 0)
}

func TestSample2(t *testing.T) {
	runSample(t, 3, 9, 0)
}

func TestSample3(t *testing.T) {
	runSample(t, 2, 4, -1)
}

func TestSample4(t *testing.T) {
	runSample(t, 5, 25, 5)
}
