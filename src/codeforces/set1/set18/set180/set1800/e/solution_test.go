package main

import "testing"

func runSample(t *testing.T, s, x string, k int, expect bool) {
	res := solve(s, x, k)

	if res != expect {
		t.Fatalf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "talant"
	x := "atltna"
	k := 3
	expect := true
	runSample(t, s, x, k, expect)
}

func TestSample2(t *testing.T) {
	s := "abacaba"
	x := "aaaabbc"
	k := 3
	expect := true
	runSample(t, s, x, k, expect)
}

func TestSample3(t *testing.T) {
	s := "abracadabraa"
	x := "avadakedavra"
	k := 3
	expect := false
	runSample(t, s, x, k, expect)
}

func TestSample4(t *testing.T) {
	s := "accio"
	x := "cicao"
	k := 3
	expect := true
	runSample(t, s, x, k, expect)
}

func TestSample5(t *testing.T) {
	s := "lumos"
	x := "molus"
	k := 3
	expect := false
	runSample(t, s, x, k, expect)
}

func TestSample6(t *testing.T) {
	s := "uwjt"
	x := "twju"
	k := 3
	expect := true
	runSample(t, s, x, k, expect)
}

func TestSample7(t *testing.T) {
	s := "kvpx"
	x := "vxpk"
	k := 3
	expect := false
	runSample(t, s, x, k, expect)
}
