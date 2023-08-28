package main

import "testing"

func runSample(t *testing.T, s string, sp string, expect int) {
	res := solve(s, sp)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "iloveslim"
	sp := "s"
	expect := 5
	runSample(t, s, sp, expect)
}

func TestSample2(t *testing.T) {
	s := "joobeel"
	sp := "oe"
	expect := 2
	runSample(t, s, sp, expect)
}

func TestSample3(t *testing.T) {
	s := "basiozi"
	sp := "si"
	expect := 3
	runSample(t, s, sp, expect)
}

func TestSample4(t *testing.T) {
	s := "khater"
	sp := "r"
	expect := 5
	runSample(t, s, sp, expect)
}

func TestSample5(t *testing.T) {
	s := "abobeih"
	sp := "abehio"
	expect := 1
	runSample(t, s, sp, expect)
}

func TestSample6(t *testing.T) {
	s := "zondl"
	sp := "abcef"
	expect := 0
	runSample(t, s, sp, expect)
}
