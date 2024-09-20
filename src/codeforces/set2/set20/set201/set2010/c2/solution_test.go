package main

import "testing"

func runSample(t *testing.T, s string, expect bool) {
	ok, res := solve(s)

	if ok != expect {
		t.Fatalf("Sample expect %t, but got %t, %s", expect, ok, res)
	}

	if !expect {
		return
	}
	n := len(s)
	m := len(res)
	if n == m {
		t.Fatalf("Sample result %s, not correct, it can't be same as %s", res, s)
	}
	// m < n
	// 中间的长度，也就是res的suf
	ln := n - m

	if ln >= m {
		t.Fatalf("Sample result %s, can't be true", res)
	}
	pref := res[:ln]
	mid := res[ln:]
	suf := res[len(mid):]

	x := pref + mid + suf

	if x != s {
		t.Fatalf("Sample result %s, getting %s, not expected %s", res, x, s)
	}
}

func TestSample1(t *testing.T) {
	s := "abrakadabrabrakadabra"
	// abrakadabra
	expect := true
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "acacacaca"
	// acacaca
	expect := true
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := "abcabc"
	expect := false
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := "abababab"
	// ababab
	expect := true
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := "tatbt"
	expect := false
	runSample(t, s, expect)
}
