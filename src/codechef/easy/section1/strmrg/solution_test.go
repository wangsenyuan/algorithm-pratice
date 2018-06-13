package main

import "testing"

func runSample(t *testing.T, n int, s1 string, m int, s2 string, expect int) {
	res := solve(n, s1, m, s2)
	if res != expect {
		t.Errorf("sample %s %s, expect %d, but got %d", s1, s2, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, m := 4, 4
	s1 := "abab"
	s2 := "baba"
	expect := 5
	runSample(t, n, s1, m, s2, expect)
}
