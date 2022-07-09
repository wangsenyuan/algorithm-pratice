package main

import "testing"

func runSample(t *testing.T, n, m int, a, b int) {
	c, d := solve(n, m)

	if a != c || b != d {
		t.Errorf("Sample expect %d/%d, but got %d/%d", a, b, c, d)
	}
}

func TestSample1(t *testing.T) {
	n := 1
	m := 10
	a, b := 1, 1
	runSample(t, n, m, a, b)
}

func TestSample2(t *testing.T) {
	n := 2
	m := 12
	a, b := 2, 3
	runSample(t, n, m, a, b)
}

func TestSample3(t *testing.T) {
	n := 100
	m := 10000
	a, b := 4301, 9525
	runSample(t, n, m, a, b)
}

func TestSample4(t *testing.T) {
	n := 3
	m := 12
	a, b := 5, 8
	runSample(t, n, m, a, b)
}

func TestSample5(t *testing.T) {
	n := 4
	m := 100
	a, b := 13, 21
	runSample(t, n, m, a, b)
}
