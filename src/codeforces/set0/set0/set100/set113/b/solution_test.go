package main

import "testing"

func runSample(t *testing.T, x string, s1 string, s2 string, expect int) {
	res := solve(x, s1, s2)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	x := "round"
	s1 := "ro"
	s2 := "ou"
	expect := 1
	runSample(t, x, s1, s2, expect)
}

func TestSample2(t *testing.T) {
	x := "codeforces"
	s1 := "code"
	s2 := "forca"
	expect := 0
	runSample(t, x, s1, s2, expect)
}

func TestSample3(t *testing.T) {
	x := "abababab"
	s1 := "a"
	s2 := "b"
	expect := 4
	runSample(t, x, s1, s2, expect)
}

func TestSample4(t *testing.T) {
	x := "aba"
	s1 := "ab"
	s2 := "ba"
	expect := 1
	runSample(t, x, s1, s2, expect)
}

func TestSample5(t *testing.T) {
	x := "aaaaaaaaaaaaaaa"
	s1 := "a"
	s2 := "a"
	expect := 15
	runSample(t, x, s1, s2, expect)
}

func TestSample6(t *testing.T) {
	x := "aaaaaaaaa"
	s1 := "aa"
	s2 := "aaa"
	expect := 7
	runSample(t, x, s1, s2, expect)
}
