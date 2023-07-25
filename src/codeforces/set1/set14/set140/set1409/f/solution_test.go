package main

import "testing"

func runSample(t *testing.T, s string, k int, x string, expect int) {
	res := solve(s, x, k)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "bbaa"
	x := "ab"
	k := 2
	expect := 3
	runSample(t, s, k, x, expect)
}

func TestSample2(t *testing.T) {
	s := "asddsaf"
	x := "sd"
	k := 3
	expect := 10
	runSample(t, s, k, x, expect)
}

func TestSample3(t *testing.T) {
	s := "qwertyhgfdsazxc"
	x := "qa"
	k := 6
	expect := 16
	runSample(t, s, k, x, expect)
}

func TestSample4(t *testing.T) {
	s := "abacaba"
	x := "aa"
	k := 2
	expect := 15
	runSample(t, s, k, x, expect)
}

func TestSample5(t *testing.T) {
	s := "cccbaacabcbbabbccbcaaabbbacacccbacaaababcabbaabacabaacbaacbaaccbcaabccabbaabcacabcaaaaaabaccbcacbbacabbaaccacabbcbcaaacaaccaccbabaaccbbccbaabbabbbaaaaacbacbcbaacbccabccbaacbacaaabbcbaaaacbccbbababacac"
	x := "ac"
	k := 1
	expect := 2464
	runSample(t, s, k, x, expect)
}
