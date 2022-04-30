package main

import "testing"

func runSample(t *testing.T, n int, m int, k int, expect int) {
	res := solve(n, m, k)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, m, k := 1, 1, 1
	expect := 1
	runSample(t, n, m, k, expect)
}


func TestSample2(t *testing.T) {
	n, m, k := 2, 2, 2
	expect := 4
	runSample(t, n, m, k, expect)
}

func TestSample3(t *testing.T) {
	n, m, k := 4, 3, 2
	expect := 45
	runSample(t, n, m, k, expect)
}

func TestSample4(t *testing.T) {
	n, m, k := 5, 7, 3
	expect := 5887
	runSample(t, n, m, k, expect)
}
