package main

import "testing"

func runSample(t *testing.T, L, R string, expect int) {
	res := solve(L, R)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	L, R := "53", "57"
	expect := 4
	runSample(t, L, R, expect)
}

func TestSample2(t *testing.T) {
	L, R := "179", "239"
	expect := 19
	runSample(t, L, R, expect)
}

func TestSample3(t *testing.T) {
	L, R := "17", "37"
	expect := 11
	runSample(t, L, R, expect)
}

func TestSample4(t *testing.T) {
	L, R := "132228", "132228"
	expect := 0
	runSample(t, L, R, expect)
}

func TestSample5(t *testing.T) {
	L, R := "54943329752812629795", "55157581939688863366"
	expect := 163
	runSample(t, L, R, expect)
}

func TestSample6(t *testing.T) {
	L, R := "88", "1914"
	expect := 28
	runSample(t, L, R, expect)
}
