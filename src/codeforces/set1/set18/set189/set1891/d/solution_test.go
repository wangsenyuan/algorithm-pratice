package main

import "testing"

func runSample(t *testing.T, l int, r int, expect int) {
	res := solve(l, r)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 4, 6, 6)
}

func TestSample2(t *testing.T) {
	runSample(t, 4, 7, 8)
}

func TestSample3(t *testing.T) {
	runSample(t, 4, 8, 9)
}

func TestSample4(t *testing.T) {
	l, r := 4, 1000
	runSample(t, l, r, bruteForce(l, r))
}

func TestSample5(t *testing.T) {
	l, r := 57, 179
	expect := 246
	runSample(t, l, r, expect)
}

func TestSample6(t *testing.T) {
	l, r := 4, 201018959
	expect := 1
	runSample(t, l, r, expect)
}

func TestSample7(t *testing.T) {
	l, r := 7, 201018960
	expect := 0
	runSample(t, l, r, expect)
}

func TestSample8(t *testing.T) {
	l, r := 179, 1000000000000000000
	expect := 41949982
	runSample(t, l, r, expect)
}
