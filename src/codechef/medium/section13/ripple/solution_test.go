package main

import "testing"

func runSample(t *testing.T, a, b string, expect string) {
	res := solve(a, b)
	if res != expect {
		t.Errorf("Sample %s %s, expect %s, but got %s", a, b, expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := "00"
	b := "11"
	expect := "100"
	runSample(t, a, b, expect)
}

func TestSample2(t *testing.T) {
	a := "0"
	b := "1"
	expect := "1"
	runSample(t, a, b, expect)
}

func TestSample3(t *testing.T) {
	a := "011"
	b := "100"
	expect := "11"
	runSample(t, a, b, expect)
}

func TestSample4(t *testing.T) {
	a := "0"
	b := "110101"
	expect := "1100110"
	runSample(t, a, b, expect)
}