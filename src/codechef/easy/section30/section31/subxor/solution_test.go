package main

import "testing"

func runSample(t *testing.T, n int, s string, expect int) {
	res := solve(n, s)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 2
	s := "10"
	expect := 3
	runSample(t, n, s, expect)
}

func TestSample2(t *testing.T) {
	n := 3
	s := "101"
	expect := 6
	runSample(t, n, s, expect)
}

func TestSample3(t *testing.T) {
	n := 4
	s := "1111"
	expect := 12
	runSample(t, n, s, expect)
}

func TestSample4(t *testing.T) {
	n := 10
	for x := 0; x < (1 << 10); x++ {
		s := toString(digitsArray(x))
		expect := sample(x)
		runSample(t, n, s, expect)
	}
}
