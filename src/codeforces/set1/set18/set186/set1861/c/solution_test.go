package main

import "testing"

func runSample(t *testing.T, s string, expect bool) {
	res := solve(s)

	if res != expect {
		t.Fatalf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "++1"
	expect := true
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "+++1--0"
	expect := false
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := "+0"
	expect := false
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := "++0-+1-+0"
	expect := true
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := "++0+-1+-0"
	expect := false
	runSample(t, s, expect)
}

func TestSample6(t *testing.T) {
	s := "+1-+0"
	expect := false
	runSample(t, s, expect)
}
