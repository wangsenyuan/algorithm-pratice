package main

import "testing"

func runSample(t *testing.T, n, k int, expect int) {
	res := solve(n, k)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, k := 2, 3
	expect := 4
	runSample(t, n, k, expect)
}

func TestSample2(t *testing.T) {
	n, k := 2, 2
	expect := 3
	runSample(t, n, k, expect)
}

func TestSample3(t *testing.T) {
	n, k := 3, 1
	expect := 1
	runSample(t, n, k, expect)
}

func TestSample4(t *testing.T) {
	n, k := 1, 3
	expect := 2
	runSample(t, n, k, expect)
}

func TestSample5(t *testing.T) {
	n, k := 1, 500
	expect := 2
	runSample(t, n, k, expect)
}

func TestSample6(t *testing.T) {
	n, k := 500, 250
	expect := 257950823
	runSample(t, n, k, expect)
}

func TestSample7(t *testing.T) {
	n, k := 1, 1
	expect := 1
	runSample(t, n, k, expect)
}
