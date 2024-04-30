package main

import "testing"

func runSample(t *testing.T, a int, b int, expect int) {
	res := solve(a, b)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := 123
	b := 222
	expect := 213
	runSample(t, a, b, expect)
}

func TestSample2(t *testing.T) {
	a := 3921
	b := 10000
	expect := 9321
	runSample(t, a, b, expect)
}

func TestSample3(t *testing.T) {
	a := 4940
	b := 5000
	expect := 4940
	runSample(t, a, b, expect)
}

func TestSample4(t *testing.T) {
	a := 8001
	b := 1009
	expect := 1008
	runSample(t, a, b, expect)
}
