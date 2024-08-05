package main

import "testing"

func runSample(t *testing.T, s1 string, s2 string, expect int) {
	res := solve(s1, s2)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s1 := "AMA"
	s2 := "ANAB"
	expect := 9
	runSample(t, s1, s2, expect)
}

func TestSample2(t *testing.T) {
	s1 := "BINUS"
	s2 := "BINANUSA"
	expect := 120
	runSample(t, s1, s2, expect)
}

func TestSample3(t *testing.T) {
	s1 := "BINUSUNIVERSITY"
	s2 := "BINANUSANTARAUNIVERSITYJAKARTA"
	expect := 151362308
	runSample(t, s1, s2, expect)
}
