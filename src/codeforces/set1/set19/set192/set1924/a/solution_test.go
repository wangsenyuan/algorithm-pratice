package main

import "testing"

func runSample(t *testing.T, n int, k int, s string, expect bool) {
	res := solve(n, k, s)

	if expect != (len(res) == 0) {
		t.Fatalf("Sample expect %t, but got %s", expect, res)
	}

	if expect {
		return
	}
	var i, j int
	for ; i < len(res); i++ {
		for j < len(s) && res[i] != s[j] {
			j++
		}
		if j == len(s) {
			break
		}
		j++
	}
	if i == len(res) {
		t.Fatalf("Sample result %s, is a subsequence of %s", res, s)
	}
}

func TestSample1(t *testing.T) {
	n, k := 2, 2
	s := "abba"
	expect := true
	runSample(t, n, k, s, expect)
}

func TestSample2(t *testing.T) {
	n, k := 2, 2
	s := "abb"
	expect := false
	runSample(t, n, k, s, expect)
}

func TestSample3(t *testing.T) {
	n, k := 3, 3
	s := "aabbccabab"
	expect := false
	runSample(t, n, k, s, expect)
}

func TestSample4(t *testing.T) {
	n, k := 4, 3
	s := "aabbccabab"
	expect := false
	runSample(t, n, k, s, expect)
}
func TestSample5(t *testing.T) {
	n, k := 3, 2
	s := "bbbaaa"
	expect := false
	runSample(t, n, k, s, expect)
}