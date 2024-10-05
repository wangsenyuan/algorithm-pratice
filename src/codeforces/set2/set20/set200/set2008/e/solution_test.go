package main

import "testing"

func runSample(t *testing.T, s string, expect int) {
	res := solve(s)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "a"
	expect := 1
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "ca"
	expect := 0
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := "aab"
	expect := 1
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := "ababa"
	expect := 1
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := "acdada"
	expect := 2
	runSample(t, s, expect)
}

func TestSample6(t *testing.T) {
	s := "ejibmyyju"
	expect := 6
	runSample(t, s, expect)
}