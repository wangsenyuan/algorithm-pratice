package main

import "testing"

func runSample(t *testing.T, n int, x int, s string, expect int) {
	res := solve(n, x, s)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, x := 7, 2
	s := "0010001"
	expect := 3
	runSample(t, n, x, s, expect)
}

func TestSample2(t *testing.T) {
	n, x := 4, 3
	s := "1010"
	expect := 1
	runSample(t, n, x, s, expect)
}

func TestSample3(t *testing.T) {
	n, x := 5, 2
	s := "00100"
	expect := 2
	runSample(t, n, x, s, expect)
}


func TestSample4(t *testing.T) {
	n, x := 5, 2
	s := "10000"
	expect := 2
	runSample(t, n, x, s, expect)
}

func TestSample5(t *testing.T) {
	n, x := 6, 2
	s := "100000"
	expect := 3
	runSample(t, n, x, s, expect)
}

func TestSample6(t *testing.T) {
	n, x := 6, 3
	s := "001000"
	expect := 2
	runSample(t, n, x, s, expect)
}