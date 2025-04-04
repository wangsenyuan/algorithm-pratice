package main

import "testing"

func runSample(t *testing.T, s string, expect int) {
	res := solve(s)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "baba"
	expect := 2
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "cc"
	expect := 0
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := "ddaa"
	expect := 3
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := "acbacddacbca"
	expect := 2
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := "eteffeteefekfektetee"
	expect := 10
	runSample(t, s, expect)
}
