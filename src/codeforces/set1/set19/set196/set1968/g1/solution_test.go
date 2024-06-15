package main

import "testing"

func runSample(t *testing.T, s string, l int, r int, expect int) {
	res := solve(s, l, r)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "aba"
	l, r := 3, 3
	expect := 0
	runSample(t, s, l, r, expect)
}

func TestSample2(t *testing.T) {
	s := "aaa"
	l, r := 3, 3
	expect := 1
	runSample(t, s, l, r, expect)
}

func TestSample3(t *testing.T) {
	s := "abacaba"
	l, r := 2, 2
	expect := 3
	runSample(t, s, l, r, expect)
}

func TestSample4(t *testing.T) {
	s := "abababcab"
	l, r := 4, 4
	expect := 2
	runSample(t, s, l, r, expect)
}

func TestSample5(t *testing.T) {
	s := "codeforces"
	l, r := 1, 1
	expect := 10
	runSample(t, s, l, r, expect)
}
func TestSample6(t *testing.T) {
	s := "abafababa"
	l, r := 3, 3
	expect := 2
	runSample(t, s, l, r, expect)
}
